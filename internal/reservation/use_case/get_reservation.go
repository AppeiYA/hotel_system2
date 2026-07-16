package reservation_usecase

import (
	"context"
	"hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
)

type GetReservation struct {
	reservationRepo reservation_ports.Repository
}

func NewGetReservation(reservationRepo reservation_ports.Repository) *GetReservation {
	return &GetReservation{reservationRepo: reservationRepo}
}

func (g *GetReservation) Execute(ctx context.Context, id string) (*domain.ReservationDetails, error) {
	return g.reservationRepo.GetReservationDetails(ctx, id)
}