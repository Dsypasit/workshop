package cloudpocket

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	db *sql.DB
}

type cloudpocket struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}

func New(db *sql.DB) *handler {

	return &handler{db}
}

func (h handler) GetCloudPockets(c echo.Context) error {
	// var err error
	// var stmt *sql.Stmt
	pockets := []cloudpocket{}

	stmt, err := h.db.Prepare("SELECT id,name FROM cloud_pocket")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		pocket := cloudpocket{}
		err = rows.Scan(&pocket.Id, &pocket.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		balance, err := h.getBalance(pocket.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		pocket.Balance = balance
		pockets = append(pockets, pocket)
	}
	return c.JSON(http.StatusOK, pockets)

}
