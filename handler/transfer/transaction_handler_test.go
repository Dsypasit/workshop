//go:buid unit

package transfer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type HandlerUtil interface {
	insertTransaction(c echo.Context) error
}

type MockHandler struct {
	transferReq   TransferRequest
	transferRes   TransferResponse
	HandlerToCall map[string]bool
}

func (a *MockHandler) InsertTransaction(c echo.Context) error {
	a.HandlerToCall["InsertTransaction"] = true
	c.Response().Status = http.StatusCreated
	err := c.Bind(&a.transferReq)
	if err != nil {
		return err
	}
	a.transferRes = TransferResponse{
		ID_transaction: 1,
		Amount:         decimal.NewFromInt(100),
		Note:           "note from req_transaction_want",
	}
	return nil
}

func (a *MockHandler) ExpectedTocall(HandlerName string) {
	if a.HandlerToCall == nil {
		a.HandlerToCall = make(map[string]bool)
	}

	a.HandlerToCall[HandlerName] = false
}

func TestTransferHandler(t *testing.T) {
	// Arange

	req_transaction_want_json := `{"sender":0,"receiver":1,"amount":"100","note":"note from req_transaction_want"}`
	res_transaction_want_json := `{"id":1,"amount":"100","note":"note from req_transaction_want"}`
	res_transaction_want := TransferResponse{
		ID_transaction: 1,
		Amount:         decimal.NewFromInt(100),
		Note:           "note from req_transaction_want",
	}

	// Act
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cloud-pockets/transaction", strings.NewReader(req_transaction_want_json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	Hmock := &MockHandler{}

	Hmock.ExpectedTocall("InsertTransaction")
	err := Hmock.InsertTransaction(c)
	res_transaction_got_json, _ := json.Marshal(Hmock.transferRes)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, c.Response().Status)
		assert.Equal(t, res_transaction_want, Hmock.transferRes)
		assert.Equal(t, string(res_transaction_want_json), string(res_transaction_got_json))
	}
}
