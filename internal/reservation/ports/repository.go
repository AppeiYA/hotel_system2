package reservation_ports

import (
	"context"
	"hotel_system2/internal/reservation/domain"
	"time"
)

type Repository interface {
	Create(
		ctx context.Context,
		reservation *domain.Reservation,
	) error

	FindByID(
		ctx context.Context,
		id string,
	) (*domain.Reservation, error)

	FindByIDForUpdate(
		ctx context.Context,
		id string,
	) (*domain.Reservation, error)

	List(
		ctx context.Context,
	) ([]domain.Reservation, error)

	ListByEmail(
		ctx context.Context,
		email string,
	) ([]domain.Reservation, error)

	Update(
		ctx context.Context,
		reservation *domain.Reservation,
	) error

	UpdateStatus(
		ctx context.Context,
		id string,
		status domain.ReservationStatus,
	) error

	HasOverlap(
		ctx context.Context,
		roomID string,
		checkIn,
		checkOut time.Time,
	) (bool, error)

	FindExpiredPending(
		ctx context.Context,
		olderThan time.Time,
	) ([]domain.Reservation, error)

	FindNoShow(
		ctx context.Context,
		before time.Time,
	) ([]domain.Reservation, error)

	GetReservationDetails(
		ctx context.Context,
		id string,
	) (*domain.ReservationDetails, error)
}