package reservation_usecase

import (
	"context"
	"time"

	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
	"hotel_system2/internal/shared/db"
)

type ExpirePending struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.Repository
}

func NewExpirePending(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.Repository,
) *ExpirePending {
	return &ExpirePending{
		txManager:       txManager,
		reservationRepo: reservationRepo,
	}
}

func (uc *ExpirePending) Execute(
	ctx context.Context,
	before time.Time,
) (int, error) {

	var expired int

	err := uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		reservations, err := uc.reservationRepo.FindExpiredPending(
			ctx,
			before,
		)
		if err != nil {
			return err
		}

		for _, reservation := range reservations {

			if err := uc.reservationRepo.UpdateStatus(
				ctx,
				reservation.ID,
				reservation_domain.ReservationStatusCancelled,
			); err != nil {
				return err
			}

			expired++
		}

		return nil
	})

	return expired, err
}