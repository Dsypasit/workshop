package cloudpocket

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h handler) GetCloudPocketById(c echo.Context) error {
	id := c.Param("id")
	pocket := cloudpocket{}
	stmt, err := h.db.Prepare("SELECT * FROM cloud_pocket WHERE id=$1")
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	row := stmt.QueryRow(intid)

	err = row.Scan(&pocket.Id, &pocket.Name, &pocket.Balance)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, pocket)

}
