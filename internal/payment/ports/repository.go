package payment_ports

import (
	"context"
	"hotel_system2/internal/payment/domain"
)

type Repository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	FindByID(ctx context.Context, id string) (*domain.Payment, error)
	FindByReference(ctx context.Context, ref string) (*domain.Payment, error)
	Update(ctx context.Context, payment *domain.Payment) error
	FindByReservationID(
    ctx context.Context,
    reservationID string,
	) (*domain.Payment, error)
}