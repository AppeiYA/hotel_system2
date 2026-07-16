package db

import (
	"context"
	"hotel_system2/internal/shared/ports"
)

type TransactionManager struct {
	db *DB
}

func NewTransactionManager(db *DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (t *TransactionManager) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = WithTx(ctx, tx)

	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

var _ ports.TransactionManagerInt = (*TransactionManager)(nil)