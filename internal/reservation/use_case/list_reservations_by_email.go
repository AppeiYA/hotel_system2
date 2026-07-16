package reservation_usecase

import (
	"context"
	"hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
)

type ListReservationByEmail struct {
	reservationRepo reservation_ports.Repository
}

func NewListReservationByEmail(reservationRepo reservation_ports.Repository) *ListReservationByEmail {
	return &ListReservationByEmail{
		reservationRepo: reservationRepo,
	}
}

func (uc *ListReservationByEmail) Execute(ctx context.Context, email string) ([]domain.Reservation, error) {
	reservations, err := uc.reservationRepo.ListByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return reservations, nil

}