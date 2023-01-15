package transfer

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

func (h *handler) getBalance(id string) (decimal.Decimal, error) {

	var err error
	var stmt *sql.Stmt
	stmt, err = h.db.Prepare("select (select coalesce(SUM(amount),0) from txn where txn.receiver = $1)-(select coalesce(SUM(amount),0) from txn where txn.sender = $1) as balance")
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	var balance decimal.Decimal
	row := stmt.QueryRow(id)

	err = row.Scan(&balance)

	if err != nil {
		return decimal.NewFromInt(0), err
	}

	return balance, nil
}
