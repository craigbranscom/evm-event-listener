package processors

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type DAOFactoryCreated struct {
	creator common.Address
}

type NFTFactoryCreated struct {
	creator common.Address
}

type ProposalFactoryCreated struct {
	creator common.Address
}

type ContractProducedEvent struct {
	contractAddress common.Address
	factoryAddress  common.Address
	producer        common.Address
}

func NewFactoryEventProcessor(logChan chan types.Log, db *sql.DB) {
	// //load contract abis
	// daoFactoryContractAbi, err := abi.JSON(strings.NewReader(string(daofactory.DaofactoryMetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// proposalFactoryContractAbi, err := abi.JSON(strings.NewReader(string(proposalfactory.ProposalfactoryMetaData.ABI)))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	//consume event from channel
	// 	vLog := <-logChan
	// 	eventHash := vLog.Topics[0]

	// 	//run logic based on event hash
	// 	switch eventHash.Hex() {
	// 	case utils.EVENT_HASH_DAO_FACTORY_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := daoFactoryContractAbi.Unpack("DAOFactoryCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := DAOFactoryCreated{
	// 			creator: unpackedEvent[0].(common.Address),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		contractStmt := fmt.Sprintf(
	// 			`INSERT INTO dao_factory_created_events (event_address, creator)
	// 			VALUES ('%s', '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.creator,
	// 		)
	// 		database.ExecuteStatement(db, contractStmt)

	// 		break
	// 	case utils.EVENT_HASH_PROPOSAL_FACTORY_CREATED:
	// 		//unpack event
	// 		unpackedEvent, err := proposalFactoryContractAbi.Unpack("ProposalFactoryCreated", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := ProposalFactoryCreated{
	// 			creator: unpackedEvent[0].(common.Address),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		contractStmt := fmt.Sprintf(
	// 			`INSERT INTO proposal_factory_created_events (event_address, creator)
	// 			VALUES ('%s', '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.creator,
	// 		)
	// 		database.ExecuteStatement(db, contractStmt)

	// 		break
	// 	case utils.EVENT_HASH_CONTRACT_PRODUCED:
	// 		//unpack event
	// 		unpackedEvent, err := daoFactoryContractAbi.Unpack("ContractProduced", vLog.Data)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		//interrogate into struct
	// 		event := ContractProducedEvent{
	// 			contractAddress: unpackedEvent[0].(common.Address),
	// 			factoryAddress:  unpackedEvent[1].(common.Address),
	// 			producer:        unpackedEvent[2].(common.Address),
	// 		}

	// 		//upsert table
	// 		trxStmt := fmt.Sprintf(
	// 			`INSERT INTO transactions VALUES ('%s', %d) ON CONFLICT DO NOTHING`,
	// 			common.HexToAddress(vLog.TxHash.Hex()),
	// 			vLog.BlockNumber,
	// 		)
	// 		database.ExecuteStatement(db, trxStmt)

	// 		//upsert table
	// 		contractStmt := fmt.Sprintf(
	// 			`INSERT INTO contract_produced_events (event_address, contract_address, factory_address, producer)
	// 			VALUES ('%s', '%s', '%s', '%s') ON CONFLICT DO NOTHING`,
	// 			vLog.Address,
	// 			event.contractAddress,
	// 			event.factoryAddress,
	// 			event.producer,
	// 		)
	// 		database.ExecuteStatement(db, contractStmt)

	// 		break
	// 	default:
	// 		fmt.Println("unrecognized event hash")
	// 		break
	// 	}
	// }
}
