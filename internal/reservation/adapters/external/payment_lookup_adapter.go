package external

import (
	"context"
	payment_domain "hotel_system2/internal/payment/domain"
	payment_ports "hotel_system2/internal/payment/ports"
)

type PaymentLookupAdapter struct {
	paymentRepo payment_ports.PaymentRepository
}

func NewPaymentLookupAdapter(repo payment_ports.PaymentRepository) *PaymentLookupAdapter {
	return &PaymentLookupAdapter{paymentRepo: repo}
}

func (a *PaymentLookupAdapter) FindPaymentIDByReservationID(ctx context.Context, reservationID string) (string, error) {
	payment, err := a.paymentRepo.FindByReservationID(ctx, reservationID)
	if err != nil {
		return "", err
	}
	return paymentIDToString(payment), nil
}

func paymentIDToString(p *payment_domain.Payment) string {
	return p.ID() // assuming payment.Payment exposes an ID() getter
}