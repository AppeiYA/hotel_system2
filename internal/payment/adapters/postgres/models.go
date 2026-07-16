package payment_postgres

import (
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

func (r paymentRow) toDomain() *payment_domain.Payment {
	return &payment_domain.Payment{
		ID:            r.ID,
		ReservationID: r.ReservationID,
		Reference:     r.Reference,
		Amount:        r.Amount,
		Status:        r.Status,
		Method:        r.Method,
		CreatedAt:     r.CreatedAt,
	}
}