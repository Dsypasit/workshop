//go:buid unit

package transfer

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
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
	trxInsert := TransferRequest{Sender: 0, Receiver: 1, Amount: decimal.NewFromInt(100), Note: "note from req_transaction_want"}
	trxWant := TransferResponse{ID_transaction: 1, Amount: decimal.NewFromInt(100), Note: "note from req_transaction_want"}
	db, mock, _ := sqlmock.New()
	row := sqlmock.NewRows([]string{"id", "amount", "note", "sender", "receiver"}).AddRow(1, 100, "note from req_transaction_want", 0, 1)
	mock.ExpectQuery("INSERT INTO txn").WithArgs(trxInsert.Amount, trxInsert.Note, trxInsert.Sender, trxInsert.Receiver).WillReturnRows(row)
	dbt := database{DB: db}

	//Act
	trxResGot := dbt.InsertTransaction(trxInsert)

	//Assert
	if assert.Nil(t, dbt.err) {
		assert.Equal(t, trxWant, trxResGot)
	}
}
