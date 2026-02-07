package processors

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransferEvent struct {
	from    common.Address
	to      common.Address
	tokenId *big.Int
}

func NewOGRE721EventProcessor(logChan chan types.Log, db *sql.DB) {
	// //load contract abi
	// contractAbi, err := abi.JSON(strings.NewReader(string(ogre721.Ogre721MetaData.ABI)))
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
	// 		//break execution if erc20 transfer
	// 		if len(vLog.Data) != 0 {
	// 			// fmt.Println("erc721 processor: erc20 transfer ignored")
	// 			break
	// 		}

	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("Transfer", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		_ = unpackedEvent

	// 		//decode
	// 		decodedTokenId, err := hexutil.Decode(vLog.Topics[3].Hex())
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		tokenId := new(big.Int)
	// 		tokenId.SetBytes(decodedTokenId)

	// 		//interrogate into struct
	// 		event := TransferEvent{
	// 			from:    common.HexToAddress(vLog.Topics[1].Hex()),
	// 			to:      common.HexToAddress(vLog.Topics[2].Hex()),
	// 			tokenId: tokenId,
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
	// 			`INSERT INTO erc721_transfer_events (event_address, from_address, to_address, token_id, trx_hash)
	// 			VALUES ('%s', '%s', '%s', %d, '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.from,
	// 			event.to,
	// 			event.tokenId,
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
