package transfer

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	// "github.com/labstack/echo/v4/middleware"
	// "github.com/lib/pq"
)

type TransferRequest struct {
	Sender   int64           `json:"sender"`
	Receiver int64           `json:"receiver"`
	Amount   decimal.Decimal `json:"amount"`
	Note     string          `json:"note"`
}

type TransferResponse struct {
	ID_transaction int64           `json:"id"`
	timestamp      time.Time       `json:"timestamp"`
	Amount         decimal.Decimal `json:"amount"`
	Note           string          `json:"note"`
}

type Err struct {
	Message string `json:"message"`
}

type Handler struct {
	Database database
}

func (h *Handler) TransferHandler(c echo.Context) error {
	transfer := TransferRequest{}
	h.Database.InitDatabase()
	defer h.Database.CloseDatabase()
	if err := c.Bind(&transfer); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	transferRes := h.Database.InsertTransaction(transfer)
	return c.JSON(http.StatusCreated, transferRes)

}
