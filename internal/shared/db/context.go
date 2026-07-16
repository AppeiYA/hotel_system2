package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txKey struct{}

// WithTx stores an active transaction on the context.
func WithTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// TxFromContext retrieves an active transaction from the context, if any.
func TxFromContext(ctx context.Context) *sqlx.Tx {
	tx, _ := ctx.Value(txKey{}).(*sqlx.Tx)
	return tx
}

// GetExecutor returns the transaction from ctx if one is active,
// otherwise falls back to the given base Executor (typically *DB).
func GetExecutor(ctx context.Context, base Executor) Executor {
	if tx := TxFromContext(ctx); tx != nil {
		return tx
	}
	return base
}