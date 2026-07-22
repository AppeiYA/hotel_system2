package external

import (
	"context"

	ledger_usecase "hotel_system2/internal/ledger/use_case"
	shareddomain "hotel_system2/internal/shared/domain"
)

type PaymentLedgerAdapter struct {
	postPaymentReceived *ledger_usecase.PostPaymentReceived
}

func NewPaymentLedgerAdapter(postPaymentReceived *ledger_usecase.PostPaymentReceived) *PaymentLedgerAdapter {
	return &PaymentLedgerAdapter{postPaymentReceived: postPaymentReceived}
}

func (a *PaymentLedgerAdapter) PostPaymentReceived(ctx context.Context, reservationID string, amount shareddomain.Money) error {
	_, err := a.postPaymentReceived.Execute(ctx, reservationID, amount)
	return err
}