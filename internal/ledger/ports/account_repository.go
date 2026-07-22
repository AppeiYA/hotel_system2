package ports

import (
	"context"
	"hotel_system2/internal/ledger/domain"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Account, error)
	FindByName(ctx context.Context, name string) (*domain.Account, error)
}