package cloudpocket

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/cloud-pockets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Basic YWRtaW46c2VjcmV0")
	rec := httptest.NewRecorder()

	newsMockPocketRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(0, "cashbox").
		AddRow(1, "test1")

	newsMockBalance := sqlmock.NewRows([]string{"balance"}).AddRow(100.23)
	newsMockBalance1 := sqlmock.NewRows([]string{"balance"}).AddRow(100.43)

	db, mock, _ := sqlmock.New()
	mock.ExpectPrepare("SELECT (.+) FROM cloud_pocket").ExpectQuery().WillReturnRows(newsMockPocketRows)
	mock.ExpectPrepare("^select[a-zA-Z0-9() ,.=$-]+SUM[a-zA-Z0-9() ,.=$-]+as balance").ExpectQuery().WithArgs(0).WillReturnRows(newsMockBalance)
	mock.ExpectPrepare("^select[a-zA-Z0-9() ,.=$-]+SUM[a-zA-Z0-9() ,.=$-]+as balance").ExpectQuery().WithArgs(1).WillReturnRows(newsMockBalance1)

	h := handler{db}
	e := echo.New()
	c := e.NewContext(req, rec)

	expected := "[{\"id\":0,\"name\":\"cashbox\",\"balance\":100.23},{\"id\":1,\"name\":\"test1\",\"balance\":100.43}]\n"

	// Act
	err := h.GetCloudPockets(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
