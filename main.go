package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Dappetizer/evm-event-listener/database"
	"github.com/Dappetizer/evm-event-listener/listener"
	"github.com/Dappetizer/evm-event-listener/processors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//load project env
	err := godotenv.Load("env.yaml")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	//dial eth client
	ethUrl := os.Getenv("ETH_WEBSOCKET_URL")
	ethClient, err := ethclient.Dial(ethUrl)
	if err != nil {
		log.Fatal("Failed to dial eth client", err)
	}

	//setup rabbit connection
	rabbitConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", err)
	}
	defer rabbitConn.Close()
	rabbitChan, err := rabbitConn.Channel()
	if err != nil {
		log.Fatal("Failed to open RabbitMQ channel", err)
	}
	defer rabbitChan.Close()

	//initialize database
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	connStr := fmt.Sprintf("postgres://%s:%s@%s?sslmode=disable", dbUser, dbPass, dbHost)
	db, err := database.NewDatabase(connStr)
	if err != nil {
		log.Fatal("Failed to initialize database", err)
	}
	_ = db

	//setup database subsystem state
	query := `INSERT INTO subsystem_state VALUES (0, 0, 0)`
	res, err := db.Client.ExecContext(context.Background(), query)
	if err != nil {
		log.Println(err)
	}
	_ = res

	//setup block listener subsystem
	blockListener, err := listener.NewBlockListener(ethClient, db)
	if err != nil {
		log.Fatal(err)
	}

	//setup block processor subsystem
	blocksBehind, err := strconv.ParseInt(os.Getenv("BLOCK_PROCESSOR_BLOCKS_BEHIND"), 10, 64)
	blockProcessor, err := processors.NewBlockProcessor(ethClient, db, rabbitChan, blocksBehind)
	if err != nil {
		log.Fatal(err)
	}

	//run subsystems
	go blockListener.RunBlockListener()
	go blockProcessor.RunBlockProcessor()

	//create processor channels and run goroutines
	// factoryProcChan := make(chan types.Log)
	// daoProcChan := make(chan types.Log)
	// ogre20ProcChan := make(chan types.Log)
	// ogre721ProcChan := make(chan types.Log)
	// proposalProcChan := make(chan types.Log)
	// bayProcChan := make(chan types.Log)
	// go processors.NewFactoryEventProcessor(factoryProcChan, db.Client)
	// go processors.NewDAOEventProcessor(daoProcChan, db.Client)
	// go processors.NewOGRE20EventProcessor(ogre20ProcChan, db.Client)
	// go processors.NewOGRE721EventProcessor(ogre721ProcChan, db.Client)
	// go processors.NewProposalEventProcessor(proposalProcChan, db.Client)
	// go processors.NewBayEventProcessor(bayProcChan, db.Client)

	//run listener goroutines
	// go listener.NewListener(client, utils.EVENT_SIG_DAO_FACTORY_CREATED, utils.EVENT_HASH_DAO_FACTORY_CREATED, factoryProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_PROPOSAL_FACTORY_CREATED, utils.EVENT_HASH_PROPOSAL_FACTORY_CREATED, factoryProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_CONTRACT_PRODUCED, utils.EVENT_HASH_CONTRACT_PRODUCED, factoryProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_MEMBER_REGISTERED, utils.EVENT_HASH_MEMBER_REGISTERED, daoProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_PROPOSAL_CREATED, utils.EVENT_HASH_PROPOSAL_CREATED, daoProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_TRANSFER, utils.EVENT_HASH_TRANSFER, ogre20ProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_TRANSFER, utils.EVENT_HASH_TRANSFER, ogre721ProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_STATUS_UPDATED, utils.EVENT_HASH_STATUS_UPDATED, proposalProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_VOTE_CAST, utils.EVENT_HASH_VOTE_CAST, proposalProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_BAY_CREATED, utils.EVENT_HASH_BAY_CREATED, bayProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_ASK_CREATED, utils.EVENT_HASH_ASK_CREATED, bayProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_BID_CREATED, utils.EVENT_HASH_BID_CREATED, bayProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_ASK_CANCELLED, utils.EVENT_HASH_ASK_CANCELLED, bayProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_BID_CANCELLED, utils.EVENT_HASH_BID_CANCELLED, bayProcChan)
	// go listener.NewListener(client, utils.EVENT_SIG_ORDER_FULFILLED, utils.EVENT_HASH_ORDER_FULFILLED, bayProcChan)

	//create executor channel and run goroutine
	// executorChan := make(chan string)
	// go executor.NewExecutor(ethUrl, executorChan)

	log.Println("Refinery startup complete")

	for {
	}
}
