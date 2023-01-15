package transfer

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionAll(t *testing.T) {

	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/transaction", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Basic YWRtaW46c2VjcmV0")
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "amount", "note", "sender", "receiver", "timestamp"}).
		AddRow(0, 100.00, "init", "0", "1", "2023-01-14T15:10:35.181Z").
		AddRow(1, 200.00, "init", "0", "1", "2023-01-14T15:15:35.181Z")

	db, mock, err := sqlmock.New()
	mock.ExpectPrepare("SELECT (.+) FROM txn").ExpectQuery().WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}
	c := e.NewContext(req, rec)

	expected := "[{\"id\":0,\"timestamp\":\"2023-01-14T15:10:35.181Z\",\"amount\":100,\"note\":\"init\",\"sender\":\"0\",\"receiver\":\"1\"},{\"id\":1,\"timestamp\":\"2023-01-14T15:15:35.181Z\",\"amount\":200,\"note\":\"init\",\"sender\":\"0\",\"receiver\":\"1\"}]\n"

	// Act
	err = h.GetTransaction(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func TestGetTransactionByPocketId(t *testing.T) {

	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/transaction/0", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Basic YWRtaW46c2VjcmV0")
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "amount", "note", "sender", "receiver", "timestamp"}).
		AddRow(0, 100.00, "init", "0", "1", "2023-01-14T15:10:35.181Z").
		AddRow(1, 200.00, "init", "0", "1", "2023-01-14T15:15:35.181Z")

	db, mock, err := sqlmock.New()
	mock.ExpectPrepare("SELECT (.+) FROM txn").ExpectQuery().WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}
	c := e.NewContext(req, rec)

	expected := "[{\"id\":0,\"timestamp\":\"2023-01-14T15:10:35.181Z\",\"amount\":100,\"note\":\"init\",\"sender\":\"0\",\"receiver\":\"1\"},{\"id\":1,\"timestamp\":\"2023-01-14T15:15:35.181Z\",\"amount\":200,\"note\":\"init\",\"sender\":\"0\",\"receiver\":\"1\"}]\n"

	// Act
	err = h.GetTransactionByPocketId(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
