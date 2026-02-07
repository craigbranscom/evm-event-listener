package listener

import (
	"context"
	"log"

	"github.com/Dappetizer/evm-event-listener/database"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockListener struct {
	ethClient *ethclient.Client
	db        *database.Database
}

func NewBlockListener(ethClient *ethclient.Client, db *database.Database) (*BlockListener, error) {
	return &BlockListener{ethClient, db}, nil
}

func (bl *BlockListener) RunBlockListener() {
	//subscribe to new block headers from chain
	headerChan := make(chan *types.Header)
	sub, err := bl.ethClient.SubscribeNewHead(context.Background(), headerChan)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	log.Println("Block Listener Subsystem: Running...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headerChan:
			log.Println("Received Header for Block", header.Number)

			//update latest header in db
			err := bl.db.SetLatestHeaderReceived(context.Background(), header.Number)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
