package ledger_postgres

import (
	"context"

	"hotel_system2/internal/ledger/domain"
)

func (r *AccountRepository) FindByID(ctx context.Context, id string) (*domain.Account, error) {
	exec := r.Executor(ctx)
	var row accountRow
	if err := exec.GetContext(ctx, &row, FindAccountByID, id); err != nil {
		return nil, err
	}
	return row.toDomain(), nil
}

func (r *AccountRepository) FindByName(ctx context.Context, name string) (*domain.Account, error) {
	exec := r.Executor(ctx)
	var row accountRow
	if err := exec.GetContext(ctx, &row, FindAccountByName, name); err != nil {
		return nil, err
	}
	return row.toDomain(), nil
}