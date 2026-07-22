package payment_usecase

import (
	"context"

	payment_domain "hotel_system2/internal/payment/domain"
	payment_ports "hotel_system2/internal/payment/ports"
	shared_ports "hotel_system2/internal/shared/ports"
)

type CompletePayment struct {
	txManager             shared_ports.TransactionManagerInt
	paymentRepo           payment_ports.PaymentRepository
	reservationConfirmer  payment_ports.ReservationConfirmationPort
	gateway               payment_ports.Gateway
	ledger                payment_ports.LedgerPort
}

func NewCompletePayment(
	txManager shared_ports.TransactionManagerInt,
	paymentRepo payment_ports.PaymentRepository,
	reservationConfirmer payment_ports.ReservationConfirmationPort,
	gateway payment_ports.Gateway,
	ledger payment_ports.LedgerPort,
) *CompletePayment {
	return &CompletePayment{
		txManager:            txManager,
		paymentRepo:          paymentRepo,
		reservationConfirmer: reservationConfirmer,
		gateway:              gateway,
		ledger:               ledger,
	}
}

func (uc *CompletePayment) Execute(ctx context.Context, reference string) error {
	payment, err := uc.paymentRepo.FindByReference(ctx, reference)
	if err != nil {
		return err
	}

	if payment.Status() == payment_domain.PaymentStatusSuccess {
		return payment_domain.ErrPaymentAlreadyCompleted
	}

	ok, method, err := uc.gateway.Verify(ctx, reference)
	if err != nil {
		return err
	}
	if !ok {
		return payment_domain.ErrPaymentFailed
	}

	return uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		payment, err := uc.paymentRepo.FindByReference(ctx, reference)
		if err != nil {
			return err
		}

		if err := payment.Complete(method); err != nil {
			return err
		}
		if err := uc.paymentRepo.Update(ctx, payment); err != nil {
			return err
		}

		if err := uc.reservationConfirmer.ConfirmReservation(ctx, payment.ReservationID()); err != nil {
			return err
		}

		if err := uc.ledger.PostPaymentReceived(ctx, payment.ReservationID(), payment.Amount()); err != nil {
			return err
		}

		return nil
	})
}