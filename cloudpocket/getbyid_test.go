package cloudpocket

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCloundPocketById(t *testing.T) {
	want := struct {
		output cloudpocket
		status int
	}{
		output: cloudpocket{
			Id:      1,
			Name:    "test1",
			Balance: 300.0,
		},
		status: http.StatusOK,
	}

	db, mock, _ := sqlmock.New()
	row := sqlmock.NewRows([]string{"id", "name", "balance"}).AddRow(1, "test1", 300.0)
	mock.ExpectPrepare("SELECT \\* FROM cloud_pocket").ExpectQuery().WithArgs(1).WillReturnRows(row)
	myhandler := handler{db}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, myhandler.GetCloudPocketById(c)) {
		assert.Equal(t, want.status, rec.Code)
		var result cloudpocket
		json.NewDecoder(rec.Body).Decode(&result)
		// fmt.Println(result)
		assert.Equal(t, want.output, result)
	}

}
