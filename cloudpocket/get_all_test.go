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
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/cloud-pockets", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Basic YWRtaW46c2VjcmV0")
	rec := httptest.NewRecorder()

	newsMockRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(0, "cashbox").
		AddRow(1, "test1")

	db, mock, err := sqlmock.New()
	mock.ExpectPrepare("SELECT (.+) FROM cloud_pocket").ExpectQuery().WillReturnRows(newsMockRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	h := handler{db}
	c := e.NewContext(req, rec)

	expected := "[{\"id\":0,\"name\":\"cashbox\",\"balance\":0},{\"id\":1,\"name\":\"test1\",\"balance\":0}]\n"

	// Act
	err = h.GetCloudPockets(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
