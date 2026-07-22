package external

import (
	"context"

	ledger_usecase "hotel_system2/internal/ledger/use_case"
	shareddomain "hotel_system2/internal/shared/domain"
)

type ReservationLedgerAdapter struct {
	postRoomCharge  *ledger_usecase.PostRoomCharge
	getFolioBalance *ledger_usecase.GetFolioBalance
}

func NewReservationLedgerAdapter(
	postRoomCharge *ledger_usecase.PostRoomCharge,
	getFolioBalance *ledger_usecase.GetFolioBalance,
) *ReservationLedgerAdapter {
	return &ReservationLedgerAdapter{postRoomCharge: postRoomCharge, getFolioBalance: getFolioBalance}
}

func (a *ReservationLedgerAdapter) PostRoomCharge(ctx context.Context, reservationID string, amount shareddomain.Money) error {
	_, err := a.postRoomCharge.Execute(ctx, reservationID, amount)
	return err
}

func (a *ReservationLedgerAdapter) GetOutstandingBalance(ctx context.Context, reservationID string) (shareddomain.Money, error) {
	return a.getFolioBalance.Execute(ctx, reservationID)
}