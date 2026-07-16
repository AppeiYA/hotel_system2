package worker

import (
	"context"
	"time"

	reservation_usecase "hotel_system2/internal/reservation/use_case"
)

type ReservationWorker struct {
	expirePending *reservation_usecase.ExpirePending
	expireNoShow *reservation_usecase.ExpireNoShow
}

func NewReservationWorker(
	expirePending *reservation_usecase.ExpirePending,
	expireNoShow *reservation_usecase.ExpireNoShow,
) *ReservationWorker {
	return &ReservationWorker{
		expirePending: expirePending,
		expireNoShow: expireNoShow,
	}
}

func (w *ReservationWorker) ExpirePendingReservations(
	ctx context.Context,
) error {

	_, err := w.expirePending.Execute(
		ctx,
		time.Now().Add(-15*time.Minute),
	)

	return err
}

func (w *ReservationWorker) ExpireNoShowReservations(
	ctx context.Context,
) error {

	_, err := w.expireNoShow.Execute(
		ctx,
		time.Now(),
	)

	return err
}