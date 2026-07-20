package reservation_postgres

import (
	"context"

	"hotel_system2/internal/reservation/domain"
)

func (r *Repository) Update(ctx context.Context, reservation *domain.Reservation) error {
	exec := r.executor(ctx)
	row := reservationRowFromDomain(reservation)

	var result reservationRow
	err := exec.QueryRowxContext(
		ctx, UpdateReservation,
		row.GuestID, row.RoomID, row.CheckIn, row.CheckOut, row.TotalAmount, row.Status, row.ID,
	).StructScan(&result)
	if err != nil {
		return err
	}

	saved, err := result.toDomain()
	if err != nil {
		return err
	}
	*reservation = *saved
	return nil
}