package database

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Database struct {
	Client *sql.DB
}

type SubsystemStateTableRow struct {
	Id                   int64
	LatestHeaderReceived int64
	LatestBlockProcessed int64
}

type BlocksTableRow struct {
	BlockNumber int64
	BlockHash   string
}

type TransactionsTableRow struct {
	TransactionHash string
	From            common.Address
	To              common.Address
	Value           *big.Int
	Data            []byte
	Status          bool
	BlockNumber     int64
}

type EventsTableRow struct {
	EventName       string
	EventSignature  string
	TransactionHash string
}

func NewDatabase(connStr string) (*Database, error) {
	//open database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	//check for liveness
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//run migrations
	path := "database/migrations.sql"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sql := string(file)
	res, err := db.Exec(sql)
	if err != nil {
		return nil, err
	}
	_ = res

	fmt.Println("Database setup complete")

	return &Database{db}, nil
}

func (db *Database) GetSubsystemState(ctx context.Context) (*SubsystemStateTableRow, error) {
	query := `SELECT * FROM subsystem_state`

	row := &SubsystemStateTableRow{}
	err := db.Client.QueryRowContext(ctx, query).Scan(&row.Id, &row.LatestHeaderReceived, &row.LatestBlockProcessed)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (db *Database) SetLatestHeaderReceived(ctx context.Context, latestHeaderReceived *big.Int) error {
	query := `UPDATE subsystem_state SET latest_header_received = $1`

	_, err := db.Client.ExecContext(ctx, query, latestHeaderReceived.Int64())
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) SetLatestBlockProcessed(ctx context.Context, latestBlockProcessed *big.Int) error {
	query := `UPDATE subsystem_state SET latest_block_processed = $1`

	_, err := db.Client.ExecContext(ctx, query, latestBlockProcessed.Int64())
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) InsertBlock(ctx context.Context, block *types.Block) (*BlocksTableRow, error) {
	query := `INSERT INTO blocks VALUES ($1, $2) RETURNING block_number, block_hash`

	row := &BlocksTableRow{}
	err := db.Client.QueryRowContext(ctx, query, block.Number().Int64(), block.Hash().Hex()).Scan(&row.BlockNumber, &row.BlockHash)
	if err != nil {
		fmt.Println("Error inserting block in database:", err)
	}

	return row, nil
}

func (db *Database) InsertTransaction(ctx context.Context, transaction *types.Transaction, blockNumber *big.Int) (*TransactionsTableRow, error) {
	query := `INSERT INTO transactions VALUES ($1, $2) RETURNING trx_hash, block_number`

	row := &TransactionsTableRow{}
	err := db.Client.QueryRowContext(ctx, query, transaction.Hash().Hex(), blockNumber.Int64()).Scan(&row.TransactionHash, &row.BlockNumber)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (db *Database) InsertEvent(ctx context.Context, eventLog *types.Log) (*EventsTableRow, error) {
	query := `INSERT INTO events VALUES (DEFAULT, $1, $2) RETURNING event_signature, trx_hash`

	row := &EventsTableRow{}
	err := db.Client.QueryRowContext(ctx, query, eventLog.Topics[0].Hex(), eventLog.TxHash.Hex()).Scan(&row.EventSignature, &row.TransactionHash)
	if err != nil {
		return nil, err
	}

	return row, nil
}
