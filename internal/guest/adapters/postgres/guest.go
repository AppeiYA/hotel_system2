package guest_postgres

import (
	"context"
	"database/sql"
	"errors"

	"hotel_system2/internal/guest/domain"
)

func (r *Repository) Create(
	ctx context.Context,
	guest *domain.Guest,
) error {

	var guestRow guestRow

	exec := r.executor(ctx)

	err := exec.QueryRowxContext(
		ctx,
		Create,
		guest.FirstName(),
		guest.LastName(),
		guest.Email().String(),
		guest.Phone(),
	).StructScan(&guestRow)

	if err != nil {
		return err
	}

	resp, err := guestRow.toDomain()
	if err != nil {
		return err
	}

	*guest = *resp

	return nil
}

func (r *Repository) FindOrCreate(
	ctx context.Context,
	guest *domain.Guest,
) error {

	var guestRow guestRow

	exec := r.executor(ctx)

	err := exec.QueryRowxContext(
		ctx,
		FindOrCreate,
		guest.FirstName,
		guest.LastName,
		guest.Email,
		guest.Phone,
	).StructScan(&guestRow)

	if err != nil {
		return err
	}

	resp, err := guestRow.toDomain()
	if err != nil {
		return err
	}
	*guest = *resp

	return nil
}

func (r *Repository) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {

	exec := r.executor(ctx)

	var exists bool

	err := exec.GetContext(
		ctx,
		&exists,
		ExistsByEmail,
		email,
	)

	return exists, err
}

func (r *Repository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.Guest, error) {

	exec := r.executor(ctx)

	var guestRow guestRow

	err := exec.GetContext(
		ctx,
		&guestRow,
		FindByEmail,
		email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrGuestNotFound
		}
		return nil, err
	}

	return guestRow.toDomain()
}

func (r *Repository) FindByID(ctx context.Context, id string) (*domain.Guest, error) {
	exec := r.executor(ctx)
	var row guestRow
	if err := exec.GetContext(ctx, &row, FindGuestByID, id); err != nil {
		return nil, err
	}
	return row.toDomain()
}