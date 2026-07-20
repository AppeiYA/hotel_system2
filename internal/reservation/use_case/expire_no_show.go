package reservation_usecase

import (
	"context"
	"time"

	reservation_ports "hotel_system2/internal/reservation/ports"
	"hotel_system2/internal/shared/db"
)

type ExpireNoShow struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.ReservationRepository
}

func NewExpireNoShow(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.ReservationRepository,
) *ExpireNoShow {
	return &ExpireNoShow{
		txManager:       txManager,
		reservationRepo: reservationRepo,
	}
}

func (uc *ExpireNoShow) Execute(
	ctx context.Context,
	before time.Time,
) (int, error) {

	var expired int

	err := uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		reservations, err := uc.reservationRepo.FindNoShow(
			ctx,
			before,
		)
		if err != nil {
			return err
		}

		for _, reservation := range reservations {
			if err := reservation.MarkNoShow(); err != nil {
				return err
			}
			if err := uc.reservationRepo.Update(ctx, reservation); err != nil {
				return err
			}
			expired++
		}

		return nil
	})

	return expired, err
}