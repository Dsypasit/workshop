package transfer

import (
	"database/sql"
	"log"

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
	createTB := `CREATE TABLE IF NOT EXISTS "txn" (
		"id" int4 NOT NULL DEFAULT nextval('txn_id'::regclass)
		"timestamp" TIMESTAMP NOT NULL,
		"amount" NUMERIC NOT NULL,
		"note" VARCHAR NOT NULL,
		"sender" int4 NOT NULL,
		"receiver" int4 NOTN NULL,
		PRIMARY KEY ("id") 
	)`

	_, db.err = db.DB.Exec(createTB)
	if db.err != nil {
		db.errMsg = db.err.Error()
		log.Fatal("cant`t create table", db.err)
	}
	log.Println("Okey Database it Have Table")
}

func (db *database) InitDatabase() {
	db.connectDatabase()
	db.createDatabase()
}

func (db *database) CloseDatabase() {
	db.DB.Close()
}
