package reservation_usecase

import (
	"context"
	guest_domain "hotel_system2/internal/guest/domain"
	guest_ports "hotel_system2/internal/guest/ports"
	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
	room_domain "hotel_system2/internal/room/domain"
	room_ports "hotel_system2/internal/room/ports"
	shared_ports "hotel_system2/internal/shared/ports"
)

type CreateReservationInput struct {
	FirstName string 
	LastName  string              
	Email     guest_domain.Email              
	Phone     string              
	RoomID    string              
	CheckIn   FlexibleDateTime 
	CheckOut  FlexibleDateTime
}

type CreateReservation struct {
	txManager shared_ports.TransactionManagerInt
	reservationRepo reservation_ports.Repository
	roomRepo room_ports.Repository
	guestRepo guest_ports.Repository
}

func NewCreateReservation(
	txManager shared_ports.TransactionManagerInt, 
	reservationRepo reservation_ports.Repository, 
	roomRepo room_ports.Repository, 
	guestRepo guest_ports.Repository,
	) *CreateReservation {
	return &CreateReservation{txManager: txManager, reservationRepo: reservationRepo, roomRepo: roomRepo, guestRepo: guestRepo}
}

func (uc *CreateReservation) Execute(
	ctx context.Context,
	input CreateReservationInput,
) (*reservation_domain.ReservationDetails, error) {

	var details *reservation_domain.ReservationDetails
	err := uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		room, err := uc.roomRepo.FindByIDForUpdate(ctx, input.RoomID)
		if err != nil {
			return err
		}

		if room.Status == room_domain.RoomStatusMaintenance {
			return room_domain.ErrRoomUnavailable
		}

		overlap, err := uc.reservationRepo.HasOverlap(
			ctx,
			room.ID,
			input.CheckIn.Time,
			input.CheckOut.Time,
		)
		if err != nil {
			return err
		}

		if overlap {
			return reservation_domain.ErrOverlappingReservation
		}

		guest := guest_domain.Guest{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Phone:     input.Phone,
		}

		if err := uc.guestRepo.FindOrCreate(ctx, &guest); err != nil {
			return err
		}

		// compute the reservation details based on the room and guest information
		totalAmount := ComputeTotalAmount(room.Rate, input.CheckIn.Time, input.CheckOut.Time)

		reservation := reservation_domain.Reservation{
			GuestID:   guest.ID,
			RoomID:    room.ID,
			CheckIn:   input.CheckIn.Time,
			CheckOut:  input.CheckOut.Time,
			TotalAmount: totalAmount,
			Status:    reservation_domain.ReservationStatusPending,
		}

		if err := uc.reservationRepo.Create(ctx, &reservation); err != nil {
			return err
		}

		details, err = uc.reservationRepo.GetReservationDetails(ctx, reservation.ID)
		
		return err
	})

	if err != nil {
		return nil, err
	}

	return details, nil
}