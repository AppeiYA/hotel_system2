package ports

import (
	"context"
	"hotel_system2/internal/ledger/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *domain.Transaction) error
	FindByID(ctx context.Context, id string) (*domain.Transaction, error)
	FindByReference(ctx context.Context, referenceType, referenceID string) ([]*domain.Transaction, error)
	MarkReversed(ctx context.Context, id string) error
}