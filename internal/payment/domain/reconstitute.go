package domain

import shared_domain "hotel_system2/internal/shared/domain"

func ReconstitutePayment(
	id string,
	reservationID string,
	reference string,
	amount int64,
	status PaymentStatus,
	method PaymentMethod,
) (*Payment, error) {
	money, err := shared_domain.NewMoney(amount)
	if err != nil {
		return nil, err
	}
	payment, err := NewPayment(id, reservationID, reference, money, method)
	if err != nil {
		return nil, err
	}
	payment.status = status
	return payment, nil
}