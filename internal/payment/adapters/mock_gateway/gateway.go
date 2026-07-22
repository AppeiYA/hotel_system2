package mock_gateway

import (
	"context"
	"fmt"

	payment_domain "hotel_system2/internal/payment/domain"
)

type Gateway struct {
	baseURL string
}

func NewGateway(baseURL string) *Gateway {
	return &Gateway{baseURL: baseURL}
}

func (g *Gateway) Initialize(ctx context.Context, email string, amount int64, reference string) (string, error) {
	return fmt.Sprintf("%s/mock-payment/%s", g.baseURL, reference), nil
}

func (g *Gateway) Verify(
	ctx context.Context,
	reference string,
) (bool, payment_domain.PaymentMethod, error) {

	return true, payment_domain.PaymentMethodCreditCard, nil
}