package reservation_postgres

import (
	"context"
	"time"

	"hotel_system2/internal/reservation/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	reservation *domain.Reservation,
) error {

	exec := r.executor(ctx)

	var row reservationRow

	err := exec.QueryRowxContext(
		ctx,
		CreateReservation,
		reservation.GuestID,
		reservation.RoomID,
		reservation.CheckIn,
		reservation.CheckOut,
		reservation.TotalAmount,
		reservation.Status,
	).StructScan(&row)
	if err != nil {
		return err
	}

	*reservation = *row.toDomain()

	return nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	id string,
) (*domain.Reservation, error) {

	exec := r.executor(ctx)

	var row reservationRow

	err := exec.GetContext(
		ctx,
		&row,
		FindByID,
		id,
	)
	if err != nil {
		return nil, err
	}

	return row.toDomain(), nil
}

func (r *Repository) FindByIDForUpdate(
	ctx context.Context,
	id string,
) (*domain.Reservation, error) {

	exec := r.executor(ctx)

	var row reservationRow

	err := exec.GetContext(
		ctx,
		&row,
		FindByIDForUpdate,
		id,
	)
	if err != nil {
		return nil, err
	}

	return row.toDomain(), nil
}

func (r *Repository) List(
	ctx context.Context,
) ([]domain.Reservation, error) {

	exec := r.executor(ctx)

	var rows []reservationRow

	err := exec.SelectContext(
		ctx,
		&rows,
		ListReservations,
	)
	if err != nil {
		return nil, err
	}

	reservations := make([]domain.Reservation, len(rows))
	for i := range rows {
		reservations[i] = *rows[i].toDomain()
	}

	return reservations, nil
}

func (r *Repository) ListByEmail(
	ctx context.Context,
	email string,
) ([]domain.Reservation, error) {

	exec := r.executor(ctx)

	var rows []reservationRow

	err := exec.SelectContext(
		ctx,
		&rows,
		ListByEmail,
		email,
	)
	if err != nil {
		return nil, err
	}

	reservations := make([]domain.Reservation, len(rows))
	for i := range rows {
		reservations[i] = *rows[i].toDomain()
	}

	return reservations, nil
}

func (r *Repository) Update(
	ctx context.Context,
	reservation *domain.Reservation,
) error {

	exec := r.executor(ctx)

	var row reservationRow

	err := exec.QueryRowxContext(
		ctx,
		UpdateReservation,
		reservation.GuestID,
		reservation.RoomID,
		reservation.CheckIn,
		reservation.CheckOut,
		reservation.TotalAmount,
		reservation.Status,
		reservation.ID,
	).StructScan(&row)
	if err != nil {
		return err
	}

	*reservation = *row.toDomain()

	return nil
}

func (r *Repository) UpdateStatus(
	ctx context.Context,
	id string,
	status domain.ReservationStatus,
) error {

	exec := r.executor(ctx)

	_, err := exec.ExecContext(
		ctx,
		UpdateStatus,
		status,
		id,
	)

	return err
}

func (r *Repository) HasOverlap(
	ctx context.Context,
	roomID string,
	checkIn,
	checkOut time.Time,
) (bool, error) {

	exec := r.executor(ctx)

	var exists bool

	err := exec.GetContext(
		ctx,
		&exists,
		HasOverlap,
		roomID,
		checkIn,
		checkOut,
	)

	return exists, err
}

func (r *Repository) FindExpiredPending(
	ctx context.Context,
	olderThan time.Time,
) ([]domain.Reservation, error) {

	exec := r.executor(ctx)

	var rows []reservationRow

	err := exec.SelectContext(
		ctx,
		&rows,
		FindExpiredPending,
		olderThan,
	)
	if err != nil {
		return nil, err
	}

	reservations := make([]domain.Reservation, len(rows))
	for i := range rows {
		reservations[i] = *rows[i].toDomain()
	}

	return reservations, nil
}

func (r *Repository) FindNoShow(
	ctx context.Context,
	before time.Time,
) ([]domain.Reservation, error) {

	exec := r.executor(ctx)

	var rows []reservationRow

	err := exec.SelectContext(
		ctx,
		&rows,
		FindNoShow,
		before,
	)
	if err != nil {
		return nil, err
	}

	reservations := make([]domain.Reservation, len(rows))
	for i := range rows {
		reservations[i] = *rows[i].toDomain()
	}

	return reservations, nil
}

func (r *Repository) GetReservationDetails(
	ctx context.Context,
	id string,
) (*domain.ReservationDetails, error) {

	exec := r.executor(ctx)

	var row reservationDetailsRow

	err := exec.GetContext(
		ctx,
		&row,
		GetReservationDetails,
		id,
	)
	if err != nil {
		return nil, err
	}

	return row.toDomain(), nil
}