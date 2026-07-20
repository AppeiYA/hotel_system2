package test_mocks

import (
	"context"

	reservation_ports "hotel_system2/internal/reservation/ports"
)

type MockPaymentLookupPort struct {
	FindPaymentIDByReservationIDFunc func(ctx context.Context, reservationID string) (string, error)
}

func (m *MockPaymentLookupPort) FindPaymentIDByReservationID(ctx context.Context, reservationID string) (string, error) {
	if m.FindPaymentIDByReservationIDFunc != nil {
		return m.FindPaymentIDByReservationIDFunc(ctx, reservationID)
	}
	return "", reservation_ports.ErrPaymentNotFound
}