package payment_ports

import (
	"context"
	shared_domain "hotel_system2/internal/shared/domain"
)

type LedgerPort interface {
	PostPaymentReceived(ctx context.Context, reservationID string, amount shared_domain.Money) error
}