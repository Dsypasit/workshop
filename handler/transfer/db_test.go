package transfer

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDatabase(t *testing.T) {
	//Arrenge
	db, mock, _ := sqlmock.New()
	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS txn`).WillReturnResult(sqlmock.NewResult(0, 0))
	dbt := database{DB: db}

	//Act
	dbt.createDatabase()

	//Assert
	assert.Nil(t, dbt.err)
}

func TestInsertTransaction(t *testing.T) {
	//Arrenge
	dtStampWant := time.Now()
	trxInsert := TransferRequest{Sender: 0, Receiver: 1, Amount: 100, Note: "note from req_transaction_want"}
	trxWant := TransferResponse{ID_transaction: 1, timestamp: dtStampWant, Amount: 100, Note: "note from req_transaction_want"}
	db, mock, _ := sqlmock.New()
	row := sqlmock.NewRows([]string{"id", "timestamp", "amount", "note", "sender", "receiver"}).AddRow(1, dtStampWant, 100, "note from req_transaction_want", 0, 1)
	mock.ExpectQuery("INSERT INTO txn").WithArgs(dtStampWant, trxInsert.Amount, trxInsert.Note, trxInsert.Sender, trxInsert.Receiver).WillReturnRows(row)
	dbt := database{DB: db}

	//Act
	trxResGot := dbt.InsertTransaction(trxInsert)

	//Assert
	if assert.Nil(t, dbt.err) {
		assert.Equal(t, trxWant, trxResGot)
	}
}
