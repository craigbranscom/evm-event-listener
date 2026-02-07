package processors

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type StatusUpdatedEvent struct {
	newStatus string
}

type VoteCastEvent struct {
	voter   common.Address
	tokenId *big.Int
	vote    uint8
}

func NewProposalEventProcessor(logChan chan types.Log, db *sql.DB) {
	// //load contract abi
	// contractAbi, err := abi.JSON(strings.NewReader(string(proposal.ProposalMetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	//consume event from channel
	// 	vLog := <-logChan
	// 	eventHash := vLog.Topics[0]

	// 	//run logic based on event hash
	// 	switch eventHash.Hex() {
	// 	case utils.EVENT_HASH_STATUS_UPDATED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("StatusUpdated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := StatusUpdatedEvent{
	// 			newStatus: unpackedEvent[0].(string),
	// 		}

	// 		//upsert transactions table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert proposals table
	// 		proposalStmt := fmt.Sprintf(
	// 			`INSERT INTO status_updated_events (event_address, new_status)
	// 			VALUES ('%s', '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.newStatus,
	// 		)
	// 		database.ExecuteStatement(db, proposalStmt)

	// 		break
	// 	case utils.EVENT_HASH_VOTE_CAST:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("VoteCast", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := VoteCastEvent{
	// 			voter:   unpackedEvent[0].(common.Address),
	// 			tokenId: unpackedEvent[1].(*big.Int),
	// 			vote:    unpackedEvent[2].(uint8),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		voteStmt := fmt.Sprintf(
	// 			`INSERT INTO vote_cast_events (event_address, voter, token_id, vote)
	// 			VALUES ('%s', '%s', %d, %d) ON CONFLICT DO NOTHING`,
	// 			vLog.Address.Hex(),
	// 			event.voter,
	// 			event.tokenId,
	// 			event.vote,
	// 		)
	// 		database.ExecuteStatement(db, voteStmt)

	// 		break
	// 	default:
	// 		fmt.Println("unrecognized event hash")
	// 		break
	// 	}
	// }
}
