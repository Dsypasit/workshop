package cloudpocket

import (
	"database/sql"
	"fmt"
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
	fmt.Println("test cloud pocket")
	var err error
	var stmt *sql.Stmt
	pockets := []cloudpocket{}

	stmt, err = h.db.Prepare("SELECT * FROM cloud_pocket")

	if err != nil {
		fmt.Println("test2")
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("test3")
		return c.JSON(http.StatusInternalServerError, err)
	}

	for rows.Next() {
		pocket := cloudpocket{}

		err = rows.Scan(&pocket.Id, &pocket.Name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		pockets = append(pockets, pocket)
	}

	return c.JSON(http.StatusOK, pockets)

}
