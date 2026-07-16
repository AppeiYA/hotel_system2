package reservation_usecase

import (
	"context"

	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_ports "hotel_system2/internal/reservation/ports"
	room_domain "hotel_system2/internal/room/domain"
	room_ports "hotel_system2/internal/room/ports"
	"hotel_system2/internal/shared/db"
)

type CheckIn struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.Repository
	roomRepo        room_ports.Repository
}

func NewCheckIn(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.Repository,
	roomRepo room_ports.Repository,
) *CheckIn {
	return &CheckIn{
		txManager:       txManager,
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
	}
}

func (uc *CheckIn) Execute(
	ctx context.Context,
	reservationID string,
) error {

	return uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {

		reservation, err := uc.reservationRepo.FindByIDForUpdate(ctx, reservationID)
		if err != nil {
			return err
		}

		if reservation.Status != reservation_domain.ReservationStatusConfirmed {
			return reservation_domain.ErrReservationNotConfirmed
		}

		room, err := uc.roomRepo.FindByIDForUpdate(ctx, reservation.RoomID)
		if err != nil {
			return err
		}

		if room.Status == room_domain.RoomStatusMaintenance {
			return room_domain.ErrRoomUnavailable
		}

		if room.Status == room_domain.RoomStatusOccupied {
			return room_domain.ErrRoomOccupied
		}

		if err := uc.roomRepo.UpdateStatus(
			ctx,
			room.ID,
			room_domain.RoomStatusOccupied,
		); err != nil {
			return err
		}

		return uc.reservationRepo.UpdateStatus(
			ctx,
			reservation.ID,
			reservation_domain.ReservationStatusCheckedIn,
		)
	})
}