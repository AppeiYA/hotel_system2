package reservation_usecase

import (
	"context"
	"hotel_system2/internal/reservation/domain"
	"hotel_system2/internal/reservation/ports"
)

type GetReservationDetailsUseCase struct {
	reservationRepo reservation_ports.ReservationRepository
	paymentLookup   reservation_ports.PaymentLookupPort
}

func NewGetReservationDetailsUseCase(
	repo reservation_ports.ReservationRepository,
	paymentLookup reservation_ports.PaymentLookupPort,
) *GetReservationDetailsUseCase {
	return &GetReservationDetailsUseCase{
		reservationRepo: repo,
		paymentLookup:   paymentLookup,
	}
}

func (uc *GetReservationDetailsUseCase) Execute(ctx context.Context, reservationID string) (*domain.ReservationDetails, error) {
	res, err := uc.reservationRepo.FindByID(ctx, reservationID)
	if err != nil {
		return nil, err
	}

	paymentID, err := uc.paymentLookup.FindPaymentIDByReservationID(ctx, reservationID)
	if err != nil {
		return nil, err
	}

	return &domain.ReservationDetails{
		Reservation: *res,
		PaymentID:   paymentID,
	}, nil
}