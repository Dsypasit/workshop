//go:build integration

package cloudpocket

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	testStmt = "INSERT INTO cloud_pocket (name, account_id) VALUES ($1, $2) RETURNING id"
)

func TestGetByID(t *testing.T) {

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}

	_ = sql.QueryRow(testStmt, "test", 1).Scan()

	h := New(sql)

	e := echo.New()
	e.GET("/cloud-pockets/:id", h.GetCloudPocketById)

	req := httptest.NewRequest(http.MethodGet, "/cloud-pockets/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	//c := e.NewContext(req, rec)
	//c.SetPath("/:id")
	//c.SetParamNames("id")
	//c.SetParamValues("1")

	//h.GetCloudPocketById(c)

	e.ServeHTTP(rec, req)

	expected := `{"id": 1, "name": "test", "balance": 0}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
