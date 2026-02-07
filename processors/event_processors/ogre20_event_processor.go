package processors

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OGRE20TransferEvent struct {
	from   common.Address
	to     common.Address
	amount *big.Int
}

func NewOGRE20EventProcessor(logChan chan types.Log, db *sql.DB) {
	// //load contract abi
	// contractAbi, err := abi.JSON(strings.NewReader(string(ogre20.Ogre20MetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	//consume event from channel
	// 	vLog := <-logChan
	// 	eventHash := vLog.Topics[0]

	// 	switch eventHash.Hex() {
	// 	case utils.EVENT_HASH_TRANSFER:
	// 		//NOTE: erc20 Transfer and erc721 Transfer events share the same hash
	// 		//break execution if erc721 transfer
	// 		if len(vLog.Data) == 0 {
	// 			// fmt.Println("erc20 processor: erc721 transfer ignored")
	// 			break
	// 		}

	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("Transfer", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := OGRE20TransferEvent{
	// 			from:   common.HexToAddress(vLog.Topics[1].Hex()),
	// 			to:     common.HexToAddress(vLog.Topics[2].Hex()),
	// 			amount: unpackedEvent[0].(*big.Int),
	// 		}

	// 		//upsert transactions table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert transfer events table
	// 		transferStmt := fmt.Sprintf(
	// 			`INSERT INTO erc20_transfer_events (event_address, from_address, to_address, amount, trx_hash)
	// 			VALUES ('%s', '%s', '%s', %d, '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.from,
	// 			event.to,
	// 			event.amount,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 		)
	// 		database.ExecuteStatement(db, transferStmt)

	// 		break
	// 	default:
	// 		fmt.Println("unrecognized event hash")
	// 		break
	// 	}
	// }
}
