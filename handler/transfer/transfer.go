package transfer

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	db *sql.DB
}

type transaction struct {
	Id        int     `json:"id"`
	Timestamp string  `json:"timestamp"`
	Amount    float32 `json:"amount"`
	Note      string  `json:"note"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

/*func (h handler) CreateTransfer(c echo.Context) error {
	exp := TransferRequest{}
	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "The request could not be found:" + err.Error()})
	}

	convtags := exp.Tags
	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", exp.Title, exp.Amount, exp.Note, pq.Array(convtags))
	err = row.Scan(&exp.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, exp)
}
*/

func (h *handler) GetTransaction(c echo.Context) error {
	txns := []transaction{}
	stmt, err := h.db.Prepare("SELECT * FROM txn ORDER BY txn.timstamp DESC")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		txn := transaction{}

		err = rows.Scan(&txn.Id, &txn.Amount, &txn.Note, &txn.Sender, &txn.Receiver, &txn.Timestamp)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		txns = append(txns, txn)
	}
	return c.JSON(http.StatusOK, txns)
}

func (h *handler) GetTransactionByPocketId(c echo.Context) error {
	id := c.Param("id")

	txns := []transaction{}
	stmt, err := h.db.Prepare("SELECT * FROM txn WHERE receiver=$1 or sender=$1 ORDER BY txn.timstamp DESC")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		txn := transaction{}

		err = rows.Scan(&txn.Id, &txn.Amount, &txn.Note, &txn.Sender, &txn.Receiver, &txn.Timestamp)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		txns = append(txns, txn)
	}
	return c.JSON(http.StatusOK, txns)
}
