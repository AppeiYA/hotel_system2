package reservation_postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"hotel_system2/internal/reservation/domain"
)

func (r *Repository) FindByID(ctx context.Context, id string) (*domain.Reservation, error) {
	exec := r.executor(ctx)
	var row reservationRow
	err := exec.GetContext(ctx, &row, FindByID, id); 
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrReservationNotFound
		}
		return nil, err
	}
	return row.toDomain()
}

func (r *Repository) List(ctx context.Context) ([]*domain.Reservation, error) {
	exec := r.executor(ctx)
	var rows []reservationRow
	if err := exec.SelectContext(ctx, &rows, ListReservations); err != nil {
		return nil, err
	}
	return mapRows(rows)
}

func (r *Repository) ListByEmail(ctx context.Context, email string) ([]*domain.Reservation, error) {
	exec := r.executor(ctx)
	var rows []reservationRow
	if err := exec.SelectContext(ctx, &rows, ListByEmail, email); err != nil {
		return nil, err
	}
	return mapRows(rows)
}

func (r *Repository) HasOverlap(ctx context.Context, roomID string, checkIn, checkOut time.Time) (bool, error) {
	exec := r.executor(ctx)
	var exists bool
	err := exec.GetContext(ctx, &exists, HasOverlap, roomID, checkIn, checkOut)
	return exists, err
}

func mapRows(rows []reservationRow) ([]*domain.Reservation, error) {
	out := make([]*domain.Reservation, 0, len(rows))
	for i := range rows {
		res, err := rows[i].toDomain()
		if err != nil {
			return nil, err
		}
		out = append(out, res)
	}
	return out, nil
}