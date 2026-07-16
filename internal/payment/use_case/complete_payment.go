package payment_usecase

import (
	"context"
	payment_domain "hotel_system2/internal/payment/domain"
	reservation_domain "hotel_system2/internal/reservation/domain"
	payment_ports "hotel_system2/internal/payment/ports"
	reservation_ports "hotel_system2/internal/reservation/ports"
	"hotel_system2/internal/shared/db"
)

type CompletePayment struct {
	txManager *db.TransactionManager

	paymentRepo payment_ports.Repository

	reservationRepo reservation_ports.Repository

	gateway payment_ports.Gateway
}

func NewCompletePayment(
	txManager *db.TransactionManager,
	paymentRepo payment_ports.Repository,
	reservationRepo reservation_ports.Repository,
	gateway payment_ports.Gateway,
) *CompletePayment {
	return &CompletePayment{
		txManager:       txManager,
		paymentRepo:     paymentRepo,
		reservationRepo: reservationRepo,
		gateway:         gateway,
	}
}


func (uc *CompletePayment) Execute(
	ctx context.Context,
	reference string,
) error {

	payment, err := uc.paymentRepo.FindByReference(ctx, reference)
	if err != nil {
		return err
	}

	if payment.Status == payment_domain.PaymentStatusSuccess {
    	return payment_domain.ErrPaymentAlreadyCompleted
	}

	ok, method, err := uc.gateway.Verify(
		ctx,
		reference,
	)
	if err != nil {
		return err
	}

	if !ok {
		return payment_domain.ErrPaymentFailed
	}

	return uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		payment, err := uc.paymentRepo.FindByReference(
			ctx,
			reference,
		)
		if err != nil {
			return err
		}

		payment.Status = payment_domain.PaymentStatusSuccess
		payment.Method = method

		if err := uc.paymentRepo.Update(ctx, payment); err != nil {
			return err
		}

		reservation, err := uc.reservationRepo.FindByIDForUpdate(
			ctx,
			payment.ReservationID,
		)
		if err != nil {
			return err
		}

		reservation.Status = reservation_domain.ReservationStatusConfirmed

		if err := uc.reservationRepo.Update(ctx, reservation); err != nil {
			return err
		}

		return nil
	})
}
