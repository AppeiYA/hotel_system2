// payment_usecase/initialize_payment.go
package payment_usecase

import (
	"context"

	payment_domain "hotel_system2/internal/payment/domain"
	payment_ports "hotel_system2/internal/payment/ports"
	guest_ports "hotel_system2/internal/guest/ports"
	reservation_ports "hotel_system2/internal/reservation/ports"

	"github.com/google/uuid"
)

type InitializePaymentInput struct {
	ReservationID string
}

type InitializePaymentResponse struct {
	Reference   string
	CheckoutURL string
}

type InitializePayment struct {
	paymentRepo     payment_ports.PaymentRepository
	reservationRepo reservation_ports.ReservationRepository
	guestRepo       guest_ports.GuestRepository
	gateway         payment_ports.Gateway
}

func NewInitializePayment(
	paymentRepo payment_ports.PaymentRepository,
	reservationRepo reservation_ports.ReservationRepository,
	guestRepo guest_ports.GuestRepository,
	gateway payment_ports.Gateway,
) *InitializePayment {
	return &InitializePayment{
		paymentRepo:     paymentRepo,
		reservationRepo: reservationRepo,
		guestRepo:       guestRepo,
		gateway:         gateway,
	}
}

func (uc *InitializePayment) Execute(ctx context.Context, input InitializePaymentInput) (*InitializePaymentResponse, error) {
	reservation, err := uc.reservationRepo.FindByID(ctx, input.ReservationID)
	if err != nil {
		return nil, err
	}

	guest, err := uc.guestRepo.FindByID(ctx, reservation.GuestID())
	if err != nil {
		return nil, err
	}

	reference := uuid.New().String()

	payment, err := payment_domain.NewPayment(
		uuid.New().String(), reservation.ID(), reference, reservation.TotalAmount(), "",
	)
	if err != nil {
		return nil, err
	}
	if err := uc.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	checkoutURL, err := uc.gateway.Initialize(ctx, guest.Email().String(), reservation.TotalAmount().AmountMinor, reference)
	if err != nil {
		return nil, err
	}

	return &InitializePaymentResponse{Reference: reference, CheckoutURL: checkoutURL}, nil
}