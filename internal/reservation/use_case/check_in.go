package reservation_usecase

import (
	"context"
	"time"

	reservation_ports "hotel_system2/internal/reservation/ports"
	room_ports "hotel_system2/internal/room/ports"
	"hotel_system2/internal/shared/db"
)

type CheckIn struct {
	txManager       *db.TransactionManager
	reservationRepo reservation_ports.ReservationRepository
	roomRepo        room_ports.RoomRepository
}

func NewCheckIn(
	txManager *db.TransactionManager,
	reservationRepo reservation_ports.ReservationRepository,
	roomRepo room_ports.RoomRepository,
) *CheckIn {
	return &CheckIn{
		txManager:       txManager,
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
	}
}

func (uc *CheckIn) Execute(ctx context.Context, reservationID string) error {
	return uc.txManager.WithinTransaction(ctx, func(ctx context.Context) error {
		reservation, err := uc.reservationRepo.FindByID(ctx, reservationID)
		if err != nil {
			return err
		}

		if err := reservation.CheckIn(time.Now()); err != nil {
			return err
		}

		room, err := uc.roomRepo.FindByID(ctx, reservation.RoomID())
		if err != nil {
			return err
		}

		if err := room.Occupy(); err != nil {
			return err
		}

		if err := uc.roomRepo.Update(ctx, room); err != nil {
			return err
		}

		return uc.reservationRepo.Update(ctx, reservation)
	})
}