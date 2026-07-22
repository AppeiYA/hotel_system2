package ports

import (
	"context"
	"hotel_system2/internal/guest/domain"
)

type GuestRepository interface {
	Create(ctx context.Context, guest *domain.Guest) error
	FindByEmail(ctx context.Context, email string) (*domain.Guest, error)
	FindByID(ctx context.Context, id string) (*domain.Guest, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindOrCreate(ctx context.Context, guest *domain.Guest) error
}