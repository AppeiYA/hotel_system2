package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
	"time"
)

type Payment struct {
	id            string
	reservationID string
	reference     string
	amount        shared_domain.Money
	method        PaymentMethod
	status        PaymentStatus
	createdAt     time.Time
}

func NewPayment(
	id, reservationID, reference string,
	amount shared_domain.Money,
	method PaymentMethod,
) (*Payment, error) {
	if reservationID == "" {
		return nil, ErrMissingReservationID
	}
	if amount.AmountMinor <= 0 {
		return nil, ErrInvalidPaymentAmount
	}
	return &Payment{
		id:            id,
		reservationID: reservationID,
		reference:     reference,
		amount:        amount,
		method:        method,
		status:        PaymentStatusPending,
		createdAt:     time.Now(),
	}, nil
}

// ---- Getters ----

func (p *Payment) ID() string                       { return p.id }
func (p *Payment) ReservationID() string             { return p.reservationID }
func (p *Payment) Reference() string                 { return p.reference }
func (p *Payment) Amount() shared_domain.Money        { return p.amount }
func (p *Payment) Method() PaymentMethod             { return p.method }
func (p *Payment) Status() PaymentStatus             { return p.status }
func (p *Payment) CreatedAt() time.Time              { return p.createdAt }

// ---- State transitions ----

func (p *Payment) Complete(method PaymentMethod) error {
	if p.status != PaymentStatusPending {
		return ErrInvalidPaymentTransition
	}
	p.status = PaymentStatusSuccess
	p.method = method
	return nil
}

func (p *Payment) Fail() error {
	if p.status != PaymentStatusPending {
		return ErrInvalidPaymentTransition
	}
	p.status = PaymentStatusFailed
	return nil
}

func (p *Payment) Refund() error {
	if p.status != PaymentStatusSuccess {
		return ErrInvalidPaymentTransition
	}
	p.status = PaymentStatusRefunded
	return nil
}