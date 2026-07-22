package reservation_postgres

import (
	"context"
	"database/sql"
	"errors"

	"hotel_system2/internal/reservation/domain"
)

func (r *Repository) Create(ctx context.Context, reservation *domain.Reservation) error {
	exec := r.executor(ctx)
	row := reservationRowFromDomain(reservation)

	var result reservationRow
	err := exec.QueryRowxContext(
		ctx, CreateReservation,
		row.GuestID, row.RoomID, row.CheckIn, row.CheckOut, row.TotalAmount, row.Status,
	).StructScan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrCreatingReservation
		}
		return err
	}

	saved, err := result.toDomain()
	if err != nil {
		return err
	}
	*reservation = *saved
	return nil
}