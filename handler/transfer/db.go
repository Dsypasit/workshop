package transfer

import (
	"database/sql"
)

type database struct {
	DB     sql.DB
	err    error
	errMsg string
}

var DB_URL string = "postgresql://group-4:group-4-pass@database-1.c7bdavepehea.ap-southeast-1.rds.amazonaws.com/group-4-dev"

/*
func (dbdatabase) connectDatabase() {
	log.Println("address database server:", DB_URL)
	db.DB, db.err = sql.Open("postgres", DB_URL)
	if db.err != nil {
		log.Fatal("Connect to database error", db.err)
	}
}

func (db *database) createDatabase() {
    createTB := CREATE TABLE IF NOT EXISTS transaction ( id SERIAL PRIMARY KEY, tstz TIMESTAMPTZ, ID_Sender INT, ID_Receiver INT, balance FLOAT  )


    , db.err = db.DB.Exec(createTB)
    if db.err != nil {
        db.errMsg = db.err.Error()
        log.Fatal("cant`t create table", db.err)
    }
    log.Println("Okey Database it Have Table")
}*/
