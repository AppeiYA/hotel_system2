package payment_ports

import (
	"context"
	reservation_domain "hotel_system2/internal/reservation/domain"
)

type ReservationConfirmationPort interface {
	ConfirmReservation(ctx context.Context, reservationID string) error
	FindReservationByID(ctx context.Context, reservationID string) (*reservation_domain.Reservation, error)
}