package external

import (
	"context"

	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
)

type ReservationConfirmationAdapter struct {
	reservationRepo reservation_ports.ReservationRepository
}

func NewReservationConfirmationAdapter(repo reservation_ports.ReservationRepository) *ReservationConfirmationAdapter {
	return &ReservationConfirmationAdapter{reservationRepo: repo}
}

func (a *ReservationConfirmationAdapter) ConfirmReservation(ctx context.Context, reservationID string) error {
	reservation, err := a.reservationRepo.FindByID(ctx, reservationID)
	if err != nil {
		return err
	}
	if err := reservation.Confirm(); err != nil {
		return err
	}
	return a.reservationRepo.Update(ctx, reservation)
}

func (a *ReservationConfirmationAdapter) FindReservationByID(ctx context.Context, reservationID string) (*reservation_domain.Reservation, error) {
	return a.reservationRepo.FindByID(ctx, reservationID)
}