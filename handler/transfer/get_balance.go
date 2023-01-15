package transfer

import (
	"database/sql"
)

func (h *handler) getBalance(id string) (float32, error) {

	var err error
	var stmt *sql.Stmt
	stmt, err = h.db.Prepare("select (select coalesce(SUM(amount),0) from txn where txn.receiver = $1)-(select coalesce(SUM(amount),0) from txn where txn.sender = $1) as balance")
	if err != nil {
		return 0, err
	}

	var balance float32
	row := stmt.QueryRow(id)

	err = row.Scan(&balance)

	if err != nil {
		return 0, err
	}

	return balance, nil
}
