package listener

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewListener(client *ethclient.Client, eventSig string, eventHash string, processorChan chan types.Log) {
	//build filter query
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{common.HexToHash(eventHash)},
		},
	}

	//create channel for event subscription
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("subscribed:", eventSig)

	loopCount := 0
	trxHashes := make(map[common.Hash]bool) //TODO: remove from map later, use hash of trxHash + trxIndex to ensure uniqueness

	//listen for events
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			loopCount++

			//only if trx hash not seen
			if trxHashes[vLog.TxHash] == false {
				fmt.Println("New Event:", eventSig, vLog.TxHash)
				trxHashes[vLog.TxHash] = true
				//send event to processor
				processorChan <- vLog //TODO: defer?
			}
		}
	}
}
