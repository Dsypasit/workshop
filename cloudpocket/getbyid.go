package cloudpocket

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCloudPocketById(c echo.Context) error {
	id := c.Param("id")
	pocket := cloudpocket{}
	stmt, err := h.db.Prepare("SELECT id,name FROM cloud_pocket WHERE id=$1")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		println("1")
		return c.JSON(http.StatusBadRequest, err)
	}
	row := stmt.QueryRow(intid)

	err = row.Scan(&pocket.Id, &pocket.Name)
	if err != nil {
		println("2")
		return c.JSON(http.StatusBadRequest, err)
	}

	balance, err := h.getBalance(pocket.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	pocket.Balance = balance
	return c.JSON(http.StatusOK, pocket)

}
