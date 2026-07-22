package reservation_ports

import (
	"context"

	shared_domain "hotel_system2/internal/shared/domain"
)

type LedgerPort interface {
	PostRoomCharge(ctx context.Context, reservationID string, amount shared_domain.Money) error
	GetOutstandingBalance(ctx context.Context, reservationID string) (shared_domain.Money, error)
}