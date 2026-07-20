package payment_postgres

import (
	"context"
	payment_domain "hotel_system2/internal/payment/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	payment *payment_domain.Payment,
) error {

	exec := r.executor(ctx)

	var row paymentRow

	err := exec.QueryRowxContext(
		ctx,
		Create,
		payment.ReservationID,
		payment.Reference,
		payment.Amount,
		payment.Status,
		payment.Method,
	).StructScan(&row)

	if err != nil {
		return err
	}

	resp, err := row.toDomain()
	if err != nil {
		return err
	}

	*payment = *resp

	return nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	id string,
) (*payment_domain.Payment, error) {

	exec := r.executor(ctx)

	var row paymentRow

	err := exec.GetContext(
		ctx,
		&row,
		FindByID,
		id,
	)

	if err != nil {
		return nil, err
	}

	resp, err := row.toDomain()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) FindByReference(
	ctx context.Context,
	reference string,
) (*payment_domain.Payment, error) {

	exec := r.executor(ctx)

	var row paymentRow

	err := exec.GetContext(
		ctx,
		&row,
		FindByReference,
		reference,
	)

	if err != nil {
		return nil, err
	}

	resp, err := row.toDomain()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *Repository) Update(
	ctx context.Context,
	payment *payment_domain.Payment,
) error {

	exec := r.executor(ctx)

	var row paymentRow

	err := exec.QueryRowxContext(
		ctx,
		Update,
		payment.Status,
		payment.Method,
		payment.ID,
	).StructScan(&row)

	if err != nil {
		return err
	}

	resp, err := row.toDomain()
	if err != nil {
		return err
	}

	*payment = *resp

	return nil
}

func (r *Repository) FindByReservationID(
    ctx context.Context,
    reservationID string,
) (*payment_domain.Payment, error) {

    exec := r.executor(ctx)

    var row paymentRow

    err := exec.GetContext(
        ctx,
        &row,
        FindByReservationID,
        reservationID,
    )
    if err != nil {
        return nil, err
    }

	resp, err := row.toDomain()
	if err != nil {
		return nil, err
	}

	return resp, nil
}