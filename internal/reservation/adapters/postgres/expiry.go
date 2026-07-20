package reservation_postgres

import (
	"context"
	"time"

	"hotel_system2/internal/reservation/domain"
)

func (r *Repository) FindExpiredPending(ctx context.Context, olderThan time.Time) ([]*domain.Reservation, error) {
	exec := r.executor(ctx)
	var rows []reservationRow
	if err := exec.SelectContext(ctx, &rows, FindExpiredPending, olderThan); err != nil {
		return nil, err
	}
	return mapRows(rows)
}

func (r *Repository) FindNoShow(ctx context.Context, before time.Time) ([]*domain.Reservation, error) {
	exec := r.executor(ctx)
	var rows []reservationRow
	if err := exec.SelectContext(ctx, &rows, FindNoShow, before); err != nil {
		return nil, err
	}
	return mapRows(rows)
}