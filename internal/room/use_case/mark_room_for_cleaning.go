package room_usecase

import (
	"context"

	"hotel_system2/internal/room/ports"
)

type MarkRoomForCleaning struct {
	roomRepo ports.RoomRepository
}

func NewMarkRoomForCleaning(roomRepo ports.RoomRepository) *MarkRoomForCleaning {
	return &MarkRoomForCleaning{roomRepo: roomRepo}
}

func (uc *MarkRoomForCleaning) Execute(ctx context.Context, roomID string) error {
	room, err := uc.roomRepo.FindByID(ctx, roomID)
	if err != nil {
		return err
	}
	if err := room.MarkForCleaning(); err != nil {
		return err
	}
	return uc.roomRepo.Update(ctx, room)
}