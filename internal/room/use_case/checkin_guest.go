package room_usecase

import (
	"context"
	reservation_ports "hotel_system2/internal/reservation/ports"
	room_ports "hotel_system2/internal/room/ports"
	"time"
)

type CheckInUseCase struct {
	roomRepo room_ports.RoomRepository
	reservationRepo reservation_ports.ReservationRepository
}

func NewCheckInUseCase(roomRepo room_ports.RoomRepository, reservationRepo reservation_ports.ReservationRepository) *CheckInUseCase {
	return &CheckInUseCase{roomRepo: roomRepo, reservationRepo: reservationRepo}
}

func (uc *CheckInUseCase) Execute(ctx context.Context, reservationID string, now time.Time) error {
	res, _ := uc.reservationRepo.FindByID(ctx, reservationID)
	if err := res.CheckIn(now); err != nil {
		return err
	}
	room, _ := uc.roomRepo.FindByID(ctx, res.RoomID()) // port into room bounded context
	if err := room.Occupy(); err != nil {
		return err
	}
	uc.reservationRepo.Update(ctx, res)
	uc.roomRepo.Update(ctx, room)
	return nil
}