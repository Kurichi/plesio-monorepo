package database

import (
	"context"
	"fmt"

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

func SetTx(ctx context.Context, tx *bun.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func (db *DB) DoInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		if r := recover(); r != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("rollback failed, err: %w", e)
			} else {
				err = fmt.Errorf("recovered from panic, err: %v", r)
			}
		}
	}()

	ctxTx := SetTx(ctx, &tx)
	if err := fn(ctxTx); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rollback failed, err: %w", err)
		}
		return err
	}

	return tx.Commit()
}
