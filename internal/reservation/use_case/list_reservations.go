package reservation_usecase

import (
	"context"
	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
)

type ListReservations struct {
	reservationRepo reservation_ports.ReservationRepository
}

func NewListReservations(reservationRepo reservation_ports.ReservationRepository) *ListReservations {
	return &ListReservations{
		reservationRepo: reservationRepo,
	}
}

func (uc *ListReservations) Execute(ctx context.Context) ([]*reservation_domain.Reservation, error) {
	reservations, err := uc.reservationRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}