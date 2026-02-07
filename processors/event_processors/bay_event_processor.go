package processors

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BayCreatedEvent struct {
	daoAddress common.Address
	admin      common.Address
}

type AskCreatedEvent struct {
	orderHash     common.Hash
	creator       common.Address
	erc721Address common.Address
	tokenId       *big.Int
	erc20Address  common.Address
	amount        *big.Int
}

type BidCreatedEvent struct {
	orderHash     common.Hash
	creator       common.Address
	erc721Address common.Address
	tokenId       *big.Int
	erc20Address  common.Address
	amount        *big.Int
}

type OrderFulfilledEvent struct {
	orderHash common.Hash
}

type AskCancelledEvent struct {
	orderHash common.Hash
}

type BidCancelledEvent struct {
	orderHash common.Hash
}

func NewBayEventProcessor(logChan chan types.Log, db *sql.DB) {
	//load contract abi
	// contractAbi, err := abi.JSON(strings.NewReader(string(bay.BayMetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	//consume event from channel
	// 	vLog := <-logChan
	// 	eventHash := vLog.Topics[0]

	// 	//run logic based on event hash
	// 	switch eventHash.Hex() {
	// 	case utils.EVENT_HASH_BAY_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("BayCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := BayCreatedEvent{
	// 			daoAddress: unpackedEvent[0].(common.Address),
	// 			admin:      unpackedEvent[1].(common.Address),
	// 		}

	// 		//upsert table
	// 		bayStmt := fmt.Sprintf(
	// 			`INSERT INTO bay_created_events (event_address, dao_address, admin_address)
	// 			VALUES ('%s', '%s', '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.daoAddress,
	// 			event.admin,
	// 		)
	// 		database.ExecuteStatement(db, bayStmt)

	// 		break
	// 	case utils.EVENT_HASH_ASK_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("AskCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := AskCreatedEvent{
	// 			orderHash:     vLog.Topics[1],
	// 			creator:       common.HexToAddress(vLog.Topics[2].Hex()),
	// 			erc721Address: common.HexToAddress(vLog.Topics[3].Hex()),
	// 			tokenId:       unpackedEvent[0].(*big.Int),
	// 			erc20Address:  unpackedEvent[1].(common.Address),
	// 			amount:        unpackedEvent[2].(*big.Int),
	// 		}

	// 		//upsert table
	// 		ordersStmt := fmt.Sprintf(
	// 			`INSERT INTO ask_created_events (event_address, order_hash, creator, erc721_address, token_id, erc20_address, amount)
	// 			VALUES ('%s', '%s', '%s', '%s', %d, '%s', %d) ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.orderHash.Hex(),
	// 			event.creator,
	// 			event.erc721Address,
	// 			event.tokenId,
	// 			event.erc20Address,
	// 			event.amount,
	// 		)
	// 		database.ExecuteStatement(db, ordersStmt)

	// 		break
	// 	case utils.EVENT_HASH_BID_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("BidCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := BidCreatedEvent{
	// 			orderHash:     vLog.Topics[1],
	// 			creator:       common.HexToAddress(vLog.Topics[2].Hex()),
	// 			erc721Address: common.HexToAddress(vLog.Topics[3].Hex()),
	// 			tokenId:       unpackedEvent[0].(*big.Int),
	// 			erc20Address:  unpackedEvent[1].(common.Address),
	// 			amount:        unpackedEvent[2].(*big.Int),
	// 		}

	// 		//upsert table
	// 		ordersStmt := fmt.Sprintf(
	// 			`INSERT INTO bid_created_events (event_address, order_hash, creator, erc721_address, token_id, erc20_address, amount)
	// 			VALUES ('%s', '%s', '%s', '%s', %d, '%s', %d) ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.orderHash.Hex(),
	// 			event.creator,
	// 			event.erc721Address,
	// 			event.tokenId,
	// 			event.erc20Address,
	// 			event.amount,
	// 		)
	// 		database.ExecuteStatement(db, ordersStmt)

	// 		break
	// 	case utils.EVENT_HASH_ASK_CANCELLED:
	// 		//interrogate into struct
	// 		event := AskCancelledEvent{
	// 			orderHash: vLog.Topics[0],
	// 		}

	// 		//upsert table
	// 		ordersStmt := fmt.Sprintf(
	// 			`INSERT INTO ask_cancelled_events (event_address, order_hash)
	// 			VALUES ('%s','%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.orderHash.Hex(),
	// 		)
	// 		database.ExecuteStatement(db, ordersStmt)

	// 		break
	// 	case utils.EVENT_HASH_BID_CANCELLED:
	// 		//interrogate into struct
	// 		event := BidCancelledEvent{
	// 			orderHash: vLog.Topics[0],
	// 		}

	// 		//upsert table
	// 		ordersStmt := fmt.Sprintf(
	// 			`INSERT INTO bid_cancelled_events (event_address, order_hash)
	// 			VALUES ('%s','%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.orderHash.Hex(),
	// 		)
	// 		database.ExecuteStatement(db, ordersStmt)

	// 		break
	// 	case utils.EVENT_HASH_ORDER_FULFILLED:
	// 		//interrogate into struct
	// 		event := OrderFulfilledEvent{
	// 			orderHash: vLog.Topics[0],
	// 		}

	// 		//upsert table
	// 		ordersStmt := fmt.Sprintf(
	// 			`INSERT INTO order_fulfilled_events (event_address, order_hash)
	// 			VALUES ('%s','%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.orderHash.Hex(),
	// 		)
	// 		database.ExecuteStatement(db, ordersStmt)

	// 		break
	// 	default:
	// 		fmt.Println("unrecognized event hash")
	// 		break
	// 	}
	// }
}
