package payment_postgres

import (
	"hotel_system2/internal/payment/domain"
	payment_domain "hotel_system2/internal/payment/domain"
	"time"
)

type paymentRow struct {
	ID            string                      `db:"id"`
	ReservationID string                      `db:"reservation_id"`
	Reference     string                      `db:"reference"`
	Amount        int64                       `db:"amount"`
	Status        payment_domain.PaymentStatus `db:"status"`
	Method        payment_domain.PaymentMethod `db:"method"`
	CreatedAt     time.Time                   `db:"created_at"`
}

func paymentRowFromDomain(payment *payment_domain.Payment) paymentRow {
	return paymentRow{
		ID:            payment.ID(),
		ReservationID: payment.ReservationID(),
		Reference:     payment.Reference(),
		Amount:        payment.Amount().AmountMinor,
		Status:        payment.Status(),
		Method:        payment.Method(),
		CreatedAt:     payment.CreatedAt(),
	}
}

func (r paymentRow) toDomain() (*payment_domain.Payment, error) {
	return domain.ReconstitutePayment(
		r.ID,
		r.ReservationID,
		r.Reference,
		r.Amount,
		r.Status,
		r.Method,
	)
}