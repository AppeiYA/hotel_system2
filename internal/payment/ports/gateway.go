package payment_ports

import (
	"context"
	payment_domain "hotel_system2/internal/payment/domain"
)

type Gateway interface {
	Initialize(
		ctx context.Context,
		email string,
		amount int64,
		reference string,
	) (authorizationURL string, err error)

	Verify(
		ctx context.Context,
		reference string,
	) (bool, payment_domain.PaymentMethod, error)
}