package account

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateAccountIT(t *testing.T) {
	// e := echo.New()

	// cfg := config.New().All()
	// sql, err := sql.Open("postgres", os.Getenv("DB_CONNECTION"))
	// if err != nil {
	// 	t.Error(err)
	// }
	// cfgFlag := config.FeatureFlag{}

	// hAccount := New(cfgFlag, sql)

	// e.POST("/accounts", hAccount.Create)

	// reqBody := `{"balance": 999.99}`
	// req := httptest.NewRequest(http.MethodPost, "/accounts", strings.NewReader(reqBody))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()

	// e.ServeHTTP(rec, req)

	// expected := `{"id": 1, "balance": 999.99}`
	// assert.Equal(t, http.StatusCreated, rec.Code)
	// assert.JSONEq(t, expected, rec.Body.String())
}
