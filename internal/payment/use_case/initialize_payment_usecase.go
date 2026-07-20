package payment_usecase

import (
	"context"
	guest_ports "hotel_system2/internal/guest/ports"
	payment_domain "hotel_system2/internal/payment/domain"
	payment_ports "hotel_system2/internal/payment/ports"
	// reservation_domain "hotel_system2/internal/reservation/domain"
	// reservation_ports "hotel_system2/internal/reservation/ports"
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

	paymentRepo     payment_ports.PaymentRepository
	reservationConfirmer  payment_ports.ReservationConfirmationPort
	guestRepo		 guest_ports.GuestRepository

	gateway payment_ports.Gateway
}

func NewInitializePayment(
	txManager *db.TransactionManager,
	paymentRepo payment_ports.PaymentRepository,
	reservationConfirmer payment_ports.ReservationConfirmationPort,
	guestRepo guest_ports.GuestRepository,
	gateway payment_ports.Gateway,
) *InitializePayment {
	return &InitializePayment{
		txManager:       txManager,
		paymentRepo:     paymentRepo,
		reservationConfirmer: reservationConfirmer,
		guestRepo: guestRepo,
		gateway:         gateway,
	}
}

func (uc *InitializePayment) Execute(
	ctx context.Context,
	input InitializePaymentInput,
) (*InitializePaymentOutput, error) {

	reservation, err := uc.reservationConfirmer.FindReservationByID(
		ctx,
		input.ReservationID,
	)
	if err != nil {
		return nil, err
	}

	reference := uuid.NewString()

	payment, err := payment_domain.NewPayment(
		uuid.New().String(),
		reservation.ID(),
		reference,
		reservation.TotalAmount(),
		payment_domain.PaymentMethodCreditCard,
	)

	if err := uc.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	existing, _ := uc.paymentRepo.FindByReservationID(
	ctx,
	reservation.ID(),
	)

	if existing != nil {
		return nil, payment_domain.ErrPaymentAlreadyInitialized
	}

	if _, err := uc.gateway.Initialize(
		ctx,
		"email",
		payment.Amount().AmountMinor,
		reference,
	); err != nil {
		return nil, err
	}

	return &InitializePaymentOutput{
		Reference: reference,
	}, nil
}