package test_mocks

import (
	"context"
	shared_domain "hotel_system2/internal/shared/domain"
)

type ReservationLedgerPorts struct {
	PostRoomChargeFunc func(ctx context.Context, reservationID string, amount shared_domain.Money) error
	GetOutstandingBalanceFunc func(ctx context.Context, reservationID string) (shared_domain.Money, error)
}

func (r *ReservationLedgerPorts) PostRoomCharge(ctx context.Context, reservationID string, amount shared_domain.Money) error {
	if r.PostRoomChargeFunc != nil {
		return r.PostRoomChargeFunc(ctx, reservationID, amount)
	}
	return nil
}

func (r *ReservationLedgerPorts) GetOutstandingBalance(ctx context.Context, reservationID string) (shared_domain.Money, error) {
	if r.GetOutstandingBalanceFunc != nil {
		return r.GetOutstandingBalanceFunc(ctx, reservationID)
	}
	return shared_domain.Money{}, nil
}