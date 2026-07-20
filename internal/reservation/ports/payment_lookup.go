package reservation_ports

import "context"

type PaymentLookupPort interface {
	FindPaymentIDByReservationID(ctx context.Context, reservationID string) (string, error)
}