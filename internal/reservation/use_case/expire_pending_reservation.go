package reservation_usecase

import (
	"context"
	"time"

	reservation_ports "hotel_system2/internal/reservation/ports"
	"hotel_system2/internal/shared/db"
)

type ExpirePending struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.ReservationRepository
}

func NewExpirePending(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.ReservationRepository,
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

			if err := reservation.MarkPending(); err != nil {
				return err
			}
			if err := uc.reservationRepo.Update(
				ctx,
				reservation,
			); err != nil {
				return err
			}

			expired++
		}

		return nil
	})

	return expired, err
}