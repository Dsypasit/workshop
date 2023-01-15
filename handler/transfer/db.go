package transfer

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type database struct {
	DB     *sql.DB
	err    error
	errMsg string
}

var DB_URL string = "postgresql://group-4:group-4-pass@database-1.c7bdavepehea.ap-southeast-1.rds.amazonaws.com/group-4-dev"

func (db *database) connectDatabase() {
	log.Println("address database server:", DB_URL)
	db.DB, db.err = sql.Open("postgres", DB_URL)
	if db.err != nil {
		log.Fatal("Connect to database error", db.err)
	}
}

func (db *database) createDatabase() {
	createTB := `CREATE TABLE IF NOT EXISTS txn ( id SERIAL PRIMARY KEY, timestamp TIMESTAMP, amount NUMERIC, note VARCHAR, sender INT, receiver INT)`

	_, db.err = db.DB.Exec(createTB)
	if db.err != nil {
		db.errMsg = db.err.Error()
		log.Fatal("cant`t create table", db.err)
	}
	log.Println("Okey Database it Have Table")
}

func (db *database) InsertTransaction(trxReq TransferRequest) TransferResponse {
	dtStamp := time.Now()
	row := db.DB.QueryRow("INSERT INTO txn (timestamp, amount, note, sender, receiver) values ($1, $2, $3, $4, $5)", dtStamp, trxReq.Amount, trxReq.Note, trxReq.Sender, trxReq.Receiver)

	resultTrxReq := TransferRequest{}
	resultTrxRes := TransferResponse{}
	db.err = row.Scan(&resultTrxRes.ID_transaction, &resultTrxRes.timestamp, &resultTrxRes.Amount, &resultTrxRes.Note, &resultTrxReq.Sender, &resultTrxReq.Receiver)
	if db.err != nil {
		log.Fatal("cant`t insert data", db.err)
		return TransferResponse{}
	}
	log.Println("insert todo success id : ", resultTrxRes.ID_transaction)
	return resultTrxRes
}

func (db *database) InitDatabase() {
	db.connectDatabase()
	db.createDatabase()
}

func (db *database) CloseDatabase() {
	db.DB.Close()
}
