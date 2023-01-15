package cloudpocket

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func TestCreateAccountIT(t *testing.T) {
	// SetupServer()

	// cp := cloudpocket{
	// 	Name: "test",
	// }
	// b, _ := json.Marshal(cp)
	// input := bytes.NewBufferString(string(b))
	// req := request("POST", "/cloud-pockets", input) //httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(reqBody))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()

	// e.ServeHTTP(rec, req)

	// req2 := httptest.NewRequest(http.MethodGet, "/clound-pockets/1", nil)
	// req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec2 := httptest.NewRecorder()

	// e.ServeHTTP(rec2, req)

	// expected := `{"id": 1, "name": "test", "balance": 0}`
	// fmt.Println(req.StatusCode)
	// fmt.Println(rec2.Body.String())
	// assert.Equal(t, http.StatusCreated, rec.Code)
	// assert.JSONEq(t, expected, rec2.Body.String())
}

var serverPort string = "2565"

func SetupServer() {
	fmt.Println("test1")
	eh := echo.New()
	go func(e *echo.Echo) {
		db, err := sql.Open("postgres", "postgresql://group-4:group-4-pass@database-1.c7bdavepehea.ap-southeast-1.rds.amazonaws.com/group-4-dev")
		if err != nil {
			fmt.Println("test2")
			log.Fatal(err)
		}
		h := New(db)
		fmt.Println("test3")
		e.POST("/cloud-pockets", h.Create)
		e.GET("/cloud-pockets/:id", h.GetCloudPocketById)

		e.Start(":" + serverPort)
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%s", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

}

func uri(paths ...string) string {
	host := "http://localhost:" + serverPort

	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	jsonStr := StreamToString(r.Body)
	return json.Unmarshal([]byte(jsonStr), v)
}
func request(method, url string, body io.Reader) *Response {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic YWRtaW46c2VjcmV0")
	res, err := http.DefaultClient.Do(req)
	return &Response{res, err}
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
