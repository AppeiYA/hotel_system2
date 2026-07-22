package ledger_postgres

import (
	"context"

	"hotel_system2/internal/ledger/domain"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *TransactionRepository) Create(ctx context.Context, tx *domain.Transaction) error {
	exec := r.Executor(ctx)

	var txRow transactionRow
	err := exec.QueryRowxContext(
		ctx, InsertTransaction,
		tx.ID(), tx.ReferenceType(), tx.ReferenceID(), tx.Description(), string(tx.Status()),
	).StructScan(&txRow)
	if err != nil {
		return err
	}

	for _, entry := range tx.Entries() {
		_, err := exec.ExecContext(
			ctx, InsertEntry,
			uuid.New().String(), tx.ID(), entry.AccountID(), string(entry.Type()), entry.Amount().AmountMinor,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *TransactionRepository) FindByID(ctx context.Context, id string) (*domain.Transaction, error) {
	exec := r.Executor(ctx)

	var txRow transactionRow
	if err := exec.GetContext(ctx, &txRow, FindTransactionByID, id); err != nil {
		return nil, err
	}

	var entryRows []entryRow
	if err := exec.SelectContext(ctx, &entryRows, FindEntriesByTransactionID, id); err != nil {
		return nil, err
	}

	return assembleTransaction(txRow, entryRows)
}

func (r *TransactionRepository) FindByReference(ctx context.Context, referenceType, referenceID string) ([]*domain.Transaction, error) {
	exec := r.Executor(ctx)

	var txRows []transactionRow
	if err := exec.SelectContext(ctx, &txRows, FindTransactionsByReference, referenceType, referenceID); err != nil {
		return nil, err
	}
	if len(txRows) == 0 {
		return nil, nil
	}

	txIDs := make([]string, len(txRows))
	for i, t := range txRows {
		txIDs[i] = t.ID
	}

	var entryRows []entryRow
	if err := exec.SelectContext(ctx, &entryRows, FindEntriesByTransactionIDs, pq.Array(txIDs)); err != nil {
		return nil, err
	}

	entriesByTx := make(map[string][]entryRow)
	for _, e := range entryRows {
		entriesByTx[e.TransactionID] = append(entriesByTx[e.TransactionID], e)
	}

	transactions := make([]*domain.Transaction, 0, len(txRows))
	for _, txRow := range txRows {
		t, err := assembleTransaction(txRow, entriesByTx[txRow.ID])
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *TransactionRepository) MarkReversed(ctx context.Context, id string) error {
	exec := r.Executor(ctx)
	_, err := exec.ExecContext(ctx, UpdateTransactionStatus, string(domain.StatusReversed), id)
	return err
}