package cloudpocket

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type PocketRequest struct {
	Name      string `json:"name"`
	AccountId int64  `json:"account_id"`
}

type PocketResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

const (
	cStmt = "INSERT INTO cloud_pocket (name, account_id) VALUES ($1, $2) RETURNING id"
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

	var id int64
	err = h.db.QueryRowContext(ctx, cStmt, req.Name, req.AccountId).Scan(&id)
	if err != nil {
		logger.Error("query row error", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "query row error", err.Error())
	}

	logger.Info("create pocket successfully", zap.Int64("id", id))

	return c.JSON(http.StatusCreated, PocketResponse{
		ID:   id,
		Name: req.Name,
	})

}

func InsertToPocketTable(tx *sql.Tx, ctx context.Context, name string, balance float64, accountId int) (int64, error) {
	var id int64
	err := tx.QueryRowContext(ctx, cStmt, name, accountId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
