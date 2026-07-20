package reservation_ports

import (
	"context"
	"time"

	"hotel_system2/internal/reservation/domain"
	custom_errors "hotel_system2/internal/shared/errors"
)

var (
	ErrPaymentNotFound = custom_errors.NotFoundError("payment not found")
)

type ReservationRepository interface {
	Create(ctx context.Context, reservation *domain.Reservation) error
	FindByID(ctx context.Context, id string) (*domain.Reservation, error)
	List(ctx context.Context) ([]*domain.Reservation, error)
	ListByEmail(ctx context.Context, email string) ([]*domain.Reservation, error)
	Update(ctx context.Context, reservation *domain.Reservation) error
	HasOverlap(ctx context.Context, roomID string, checkIn, checkOut time.Time) (bool, error)
	FindExpiredPending(ctx context.Context, olderThan time.Time) ([]*domain.Reservation, error)
	FindNoShow(ctx context.Context, before time.Time) ([]*domain.Reservation, error)
}