package processors

import (
	"context"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/Dappetizer/evm-event-listener/database"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BlockProcessor struct {
	ethClient    *ethclient.Client
	db           *database.Database
	rabbit       *amqp.Channel
	blocksBehind int64
}

func NewBlockProcessor(ethClient *ethclient.Client, db *database.Database, rabbit *amqp.Channel, blocksBehind int64) (*BlockProcessor, error) {
	return &BlockProcessor{ethClient, db, rabbit, blocksBehind}, nil
}

func (bp *BlockProcessor) RunBlockProcessor() {
	//declare rabbit exchanges
	err := bp.rabbit.ExchangeDeclare(
		"blocks", // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare rabbit blocks exchange", err)
	}
	err = bp.rabbit.ExchangeDeclare(
		"transactions", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare rabbit transactions exchange", err)
	}
	err = bp.rabbit.ExchangeDeclare(
		"events", // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare rabbit events exchange", err)
	}

	log.Println("Block Processor Subsystem: Running...")

	for {
		log.Println("Block Processor Checking for Available Work...")

		//get subsystem state
		row, err := bp.db.GetSubsystemState(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if row.LatestBlockProcessed < row.LatestHeaderReceived-bp.blocksBehind {
			blockToProcess := big.NewInt(row.LatestBlockProcessed)
			blockToProcess.Add(blockToProcess, big.NewInt(1))

			//query chain for block
			block, err := bp.ethClient.BlockByNumber(context.Background(), blockToProcess)
			if err != nil {
				log.Fatal(err)
			}

			//process block
			processBlockErr := bp.processBlock(context.Background(), block)
			if processBlockErr != nil {
				log.Fatal("Failed to process block:", processBlockErr)
			}
		} else if row.LatestBlockProcessed == row.LatestHeaderReceived-bp.blocksBehind {
			log.Println("Block Processor Caught Up")
		} else {
			log.Println("Block Processor Awaiting Work...")
		}

		processorIntervalMs, err := strconv.ParseInt(os.Getenv("BLOCK_PROCESSOR_INTERVAL_MILLISECONDS"), 10, 64)
		time.Sleep(time.Duration(processorIntervalMs) * time.Millisecond)
	}
}

func (bp *BlockProcessor) processBlock(ctx context.Context, block *types.Block) error {
	log.Println("Processing Block", block.Number())

	//insert db row
	row, err := bp.db.InsertBlock(ctx, block)
	if err != nil {
		return err
	}
	_ = row

	//process each transaction in block
	log.Println(len(block.Transactions()), "Transactions in Block", block.Number())
	for _, transaction := range block.Transactions() {
		err := bp.processTransaction(ctx, transaction, block.Number())
		if err != nil {
			return err
		}
	}

	bp.db.SetLatestBlockProcessed(ctx, block.Number())

	//build rabbit msg
	// body := block.Number().Bytes()

	//publish to rabbit
	// err = bp.rabbit.PublishWithContext(ctx,
	// 	"blocks", // exchange
	// 	"number", // routing key
	// 	false,    // mandatory
	// 	false,    // immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        body,
	// 	})
	// if err != nil {
	// 	return err
	// }

	log.Println("Block", block.Number(), "published to Rabbit")

	return nil
}

func (bp *BlockProcessor) processTransaction(ctx context.Context, transaction *types.Transaction, blockNumber *big.Int) error {
	log.Println("Processing Transaction:", transaction.Hash().Hex())

	//insert db row
	row, err := bp.db.InsertTransaction(ctx, transaction, blockNumber)
	if err != nil {
		return err
	}
	_ = row

	//get trx receipt from chain
	receipt, err := bp.ethClient.TransactionReceipt(ctx, transaction.Hash())
	if err != nil {
		return err
	}

	//process events
	log.Println(len(receipt.Logs), "Events in Transaction", transaction.Hash().Hex())
	for _, log := range receipt.Logs {
		err := bp.processEvent(ctx, log)
		if err != nil {
			return err
		}
	}

	//build rabbit msg
	// routingKey := ""
	// if receipt.Status == uint64(0) {
	// 	routingKey = "successful"
	// } else {
	// 	routingKey = "failed"
	// }
	// body := []byte(transaction.Hash().Hex())

	//publish to rabbit
	// err = bp.rabbit.PublishWithContext(ctx,
	// 	"transactions", // exchange
	// 	routingKey,     // routing key
	// 	false,          // mandatory
	// 	false,          // immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        body,
	// 	})
	// if err != nil {
	// 	return err
	// }

	// log.Println("Transaction", transaction.Hash().Hex(), "published to Rabbit")

	return nil
}

func (bp *BlockProcessor) processEvent(ctx context.Context, eventLog *types.Log) error {
	log.Println("Processing Event:", eventLog.Topics[0].Hex())

	//insert db row
	row, err := bp.db.InsertEvent(ctx, eventLog)
	if err != nil {
		return err
	}
	_ = row

	//TODO: decode event data

	//build rabbit msg
	body := []byte(eventLog.Topics[0].Hex())

	//publish to rabbit
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	err = bp.rabbit.PublishWithContext(ctx,
		"events", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}
	log.Println("Event", eventLog.Topics[0].Hex(), "published to Rabbit")

	return nil
}
