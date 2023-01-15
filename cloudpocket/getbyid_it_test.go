//go:build integration

package cloudpocket

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountIT(t *testing.T) {
	e := echo.New()

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}
	cfgFlag := config.FeatureFlag{}

	h = handler{&sql}
	reqBody := `{"balance": 999.99}`

	e.GET("/cloud-pockets/:id", h.GetCloundPocketById)
	e.POST("/cloud-pockets", h.Creat)

	reqBody := `{"name": "test"}`
	req := httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	req := httptest.NewRequest(http.MethodPost, "/clound-pockets/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	expected := `{"id": 1, "balance": 999.99}`
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
