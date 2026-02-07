package processors

import (
	"database/sql"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MemberRegisteredEvent struct {
	daoAddress    common.Address
	nftAddress    common.Address
	tokenId       *big.Int
	memberAddress common.Address
}

type ProposalCreatedEvent struct {
	daoAddress common.Address
	proposal   common.Address
	proposalId *big.Int
	creator    common.Address
}

func NewDAOEventProcessor(logChan chan types.Log, db *sql.DB) {
	// //load contract abi
	// contractAbi, err := abi.JSON(strings.NewReader(string(dao.DaoMetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	//consume event from channel
	// 	vLog := <-logChan
	// 	eventHash := vLog.Topics[0]

	// 	//run logic based on event hash
	// 	switch eventHash.Hex() {
	// 	case utils.EVENT_HASH_MEMBER_REGISTERED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("MemberRegistered", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := MemberRegisteredEvent{
	// 			daoAddress:    unpackedEvent[0].(common.Address),
	// 			nftAddress:    unpackedEvent[1].(common.Address),
	// 			tokenId:       unpackedEvent[2].(*big.Int),
	// 			memberAddress: unpackedEvent[3].(common.Address),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		memberStmt := fmt.Sprintf(
	// 			`INSERT INTO member_registered_events (event_address, dao_address, nft_address, token_id, member_address)
	// 			VALUES ('%s', '%s', '%s', %d, '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.daoAddress,
	// 			event.nftAddress,
	// 			event.tokenId,
	// 			event.memberAddress,
	// 		)
	// 		database.ExecuteStatement(db, memberStmt)

	// 		break
	// 	case utils.EVENT_HASH_PROPOSAL_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := contractAbi.Unpack("ProposalCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := ProposalCreatedEvent{
	// 			daoAddress: unpackedEvent[0].(common.Address),
	// 			proposal:   unpackedEvent[1].(common.Address),
	// 			proposalId: unpackedEvent[2].(*big.Int),
	// 			creator:    unpackedEvent[3].(common.Address),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		proposalStmt := fmt.Sprintf(
	// 			`INSERT INTO proposal_created_events (event_address, dao_address, proposal_address, proposal_id, creator)
	// 			VALUES ('%s', '%s', '%s', %d, '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.daoAddress,
	// 			event.proposal,
	// 			event.proposalId,
	// 			event.creator,
	// 		)
	// 		database.ExecuteStatement(db, proposalStmt)

	// 		break
	// 	default:
	// 		fmt.Println("unrecognized event hash")
	// 		break
	// 	}
	// }
}
