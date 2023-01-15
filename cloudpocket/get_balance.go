package cloudpocket

import (
	"database/sql"
)

func (h *handler) getBalance(id int) (float32, error) {

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

	// fmt.Print("id : ")
	// fmt.Println(id)
	// fmt.Print("balance : ")
	// fmt.Println(balance)

	// fmt.Println("")
	// fmt.Println("")

	return balance, nil
}
