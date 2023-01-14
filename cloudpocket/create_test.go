//go:buid unit

package cloudpocket

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePocket(t *testing.T) {
	tests := []struct {
		name       string
		sqlFn      func() (*sql.DB, error)
		reqBody    string
		wantStatus int
		wantBody   string
	}{
		{"create pocket succesfully",
			func() (*sql.DB, error) {
				db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				if err != nil {
					return nil, err
				}
				row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "pocket-name")
				mock.ExpectQuery(cStmt).WithArgs("pocket-name").WillReturnRows(row)
				return db, err
			},
			`{"name": "pocket-name"}`,
			http.StatusCreated,
			`{"id": 1, "name": "pocket-name", "balance": 0}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(tc.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			db, err := tc.sqlFn()
			h := New(db)
			// Assertions
			assert.NoError(t, err)
			if assert.NoError(t, h.Create(c)) {
				assert.Equal(t, tc.wantStatus, rec.Code)
				assert.JSONEq(t, tc.wantBody, rec.Body.String())
			}
		})
	}
}
