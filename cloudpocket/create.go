package cloudpocket

import (
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type PocketRequest struct {
	Name string `json:"name"`
}

type PocketResponse struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}

const (
	cStmt = "INSERT INTO cloud_pocket (name) VALUES ($1) RETURNING id, name"
)

func (h handler) Create(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()
	var req PocketRequest
	err := c.Bind(&req)
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "bad request body", err.Error())
	}

	var resp PocketResponse
	err = h.db.QueryRowContext(ctx, cStmt, req.Name).Scan(&resp.ID, &resp.Name)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "query row error", err.Error())
	}

	logger.Info("create pocket successfully", zap.Int64("id", resp.ID))

	resp.Balance = 0

	return c.JSON(http.StatusCreated, resp)

}
