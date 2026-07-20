package reservation_usecase

import (
	"context"
	"errors"
	"fmt"

	guest_domain "hotel_system2/internal/guest/domain"
	guest_ports "hotel_system2/internal/guest/ports"
	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
	room_domain "hotel_system2/internal/room/domain"
	room_ports "hotel_system2/internal/room/ports"
	shared_domain "hotel_system2/internal/shared/domain"
	shared_ports "hotel_system2/internal/shared/ports"

	"github.com/google/uuid"
)

type CreateReservationInput struct {
	FirstName string
	LastName  string
	Email     shared_domain.Email
	Phone     string
	RoomID    string
	CheckIn   FlexibleDateTime
	CheckOut  FlexibleDateTime
}

type CreateReservation struct {
	txManager       shared_ports.TransactionManagerInt
	reservationRepo reservation_ports.ReservationRepository
	roomRepo        room_ports.RoomRepository
	guestRepo       guest_ports.GuestRepository
	paymentLookup   reservation_ports.PaymentLookupPort
}

func NewCreateReservation(
	txManager shared_ports.TransactionManagerInt,
	reservationRepo reservation_ports.ReservationRepository,
	roomRepo room_ports.RoomRepository,
	guestRepo guest_ports.GuestRepository,
	paymentLookup reservation_ports.PaymentLookupPort,
) *CreateReservation {
	return &CreateReservation{
		txManager:       txManager,
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
		guestRepo:       guestRepo,
		paymentLookup:   paymentLookup,
	}
}

func (uc *CreateReservation) Execute(
	ctx context.Context,
	input CreateReservationInput,
) (*reservation_domain.ReservationDetails, error) {

	var details *reservation_domain.ReservationDetails

	err := uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		room, err := uc.roomRepo.FindByID(ctx, input.RoomID)
		if err != nil {
			return err
		}

		if room.Status() == room_domain.RoomStatusMaintenance {
			return room_domain.ErrRoomUnavailable
		}

		overlap, err := uc.reservationRepo.HasOverlap(ctx, room.ID(), input.CheckIn.Time, input.CheckOut.Time)
		if err != nil {
			return err
		}
		if overlap {
			return reservation_domain.ErrOverlappingReservation
		}

		guest, err := uc.guestRepo.FindByEmail(ctx, input.Email.String())
		if err != nil {
			if !errors.Is(err, guest_domain.ErrGuestNotFound) {
				return err
			}
			guest, err = guest_domain.NewGuest(
				uuid.New().String(),
				input.FirstName,
				input.LastName,
				input.Email,
				input.Phone,
			)
			if err != nil {
				return err
			}
			if err := uc.guestRepo.Create(ctx, guest); err != nil {
				return err
			}
		}
		fmt.Printf("Guest ID: %v\n", guest.ID())

		dateRange, err := reservation_domain.NewDateRange(input.CheckIn.Time, input.CheckOut.Time)
		if err != nil {
			return err
		}

		totalAmount, err := room.Rate().MultiplyByNights(dateRange.Nights())
		if err != nil {
			return err
		}

		fmt.Printf("Total Amount: %v\n", totalAmount)

		reservation, err := reservation_domain.NewReservation(
			uuid.New().String(),
			guest.ID(),
			room.ID(),
			dateRange,
			totalAmount,
		)
		if err != nil {
			return err
		}

		if err := uc.reservationRepo.Create(ctx, reservation); err != nil {
			return err
		}

		details, err = uc.buildDetails(ctx, reservation)
		return err
	})

	if err != nil {
		return nil, err
	}
	return details, nil
}

func (uc *CreateReservation) buildDetails(ctx context.Context, reservation *reservation_domain.Reservation) (*reservation_domain.ReservationDetails, error) {
	paymentID, err := uc.paymentLookup.FindPaymentIDByReservationID(ctx, reservation.ID())
	if err != nil && !errors.Is(err, reservation_ports.ErrPaymentNotFound) {
		return nil, err
	}

	return &reservation_domain.ReservationDetails{
		Reservation: *reservation,
		PaymentID:   paymentID,
	}, nil
}