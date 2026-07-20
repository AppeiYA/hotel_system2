package room_usecase

import (
	"context"

	"hotel_system2/internal/room/ports"
)

type OccupyRoom struct {
	roomRepo ports.RoomRepository
}

func NewOccupyRoom(roomRepo ports.RoomRepository) *OccupyRoom {
	return &OccupyRoom{roomRepo: roomRepo}
}

func (uc *OccupyRoom) Execute(ctx context.Context, roomID string) error {
	room, err := uc.roomRepo.FindByID(ctx, roomID)
	if err != nil {
		return err
	}
	if err := room.Occupy(); err != nil {
		return err
	}
	return uc.roomRepo.Update(ctx, room)
}