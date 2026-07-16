package payment_usecase

import (
	"context"
	guest_ports "hotel_system2/internal/guest/ports"
	payment_domain "hotel_system2/internal/payment/domain"
	payment_ports "hotel_system2/internal/payment/ports"
	// reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
	"hotel_system2/internal/shared/db"

	"github.com/google/uuid"
)

type InitializePaymentInput struct {
	ReservationID string
}

type InitializePaymentOutput struct {
	Reference string
}

type InitializePayment struct {
	txManager *db.TransactionManager

	paymentRepo     payment_ports.Repository
	reservationRepo reservation_ports.Repository
	guestRepo		 guest_ports.Repository

	gateway payment_ports.Gateway
}

func NewInitializePayment(
	txManager *db.TransactionManager,
	paymentRepo payment_ports.Repository,
	reservationRepo reservation_ports.Repository,
	guestRepo guest_ports.Repository,
	gateway payment_ports.Gateway,
) *InitializePayment {
	return &InitializePayment{
		txManager:       txManager,
		paymentRepo:     paymentRepo,
		reservationRepo: reservationRepo,
		guestRepo: guestRepo,
		gateway:         gateway,
	}
}

func (uc *InitializePayment) Execute(
	ctx context.Context,
	input InitializePaymentInput,
) (*InitializePaymentOutput, error) {

	reservation, err := uc.reservationRepo.FindByID(
		ctx,
		input.ReservationID,
	)
	if err != nil {
		return nil, err
	}

	reference := uuid.NewString()

	payment := payment_domain.Payment{
		ReservationID: reservation.ID,
		Reference: reference,
		Amount: reservation.TotalAmount,
		Status: payment_domain.PaymentStatusPending,
	}

	if err := uc.paymentRepo.Create(ctx, &payment); err != nil {
		return nil, err
	}

	if _, err := uc.gateway.Initialize(
		ctx,
		"email",
		payment.Amount,
		reference,
	); err != nil {
		return nil, err
	}

	return &InitializePaymentOutput{
		Reference: reference,
	}, nil
}