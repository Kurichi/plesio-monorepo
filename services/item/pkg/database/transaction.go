package database

import (
	"context"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

var txKey = struct{}{}

func GetTxByContext(ctx context.Context) (bun.Tx, error) {
	if tx, ok := ctx.Value(txKey).(bun.Tx); ok {
		return tx, nil
	}
	return bun.Tx{}, errors.New("tx not found in context")
}
