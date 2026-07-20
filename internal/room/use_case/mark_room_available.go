package room_usecase

import (
	"context"

	"hotel_system2/internal/room/ports"
)

type MarkRoomAvailable struct {
	roomRepo ports.RoomRepository
}

func NewMarkRoomAvailable(roomRepo ports.RoomRepository) *MarkRoomAvailable {
	return &MarkRoomAvailable{roomRepo: roomRepo}
}

func (uc *MarkRoomAvailable) Execute(ctx context.Context, roomID string) error {
	room, err := uc.roomRepo.FindByID(ctx, roomID)
	if err != nil {
		return err
	}
	if err := room.MarkAvailable(); err != nil {
		return err
	}
	return uc.roomRepo.Update(ctx, room)
}