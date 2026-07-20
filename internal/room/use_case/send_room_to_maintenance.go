package room_usecase

import (
	"context"

	"hotel_system2/internal/room/ports"
)

type SendRoomToMaintenance struct {
	roomRepo ports.RoomRepository
}

func NewSendRoomToMaintenance(roomRepo ports.RoomRepository) *SendRoomToMaintenance {
	return &SendRoomToMaintenance{roomRepo: roomRepo}
}

func (uc *SendRoomToMaintenance) Execute(ctx context.Context, roomID string) error {
	room, err := uc.roomRepo.FindByID(ctx, roomID)
	if err != nil {
		return err
	}
	if err := room.SendToMaintenance(); err != nil {
		return err
	}
	return uc.roomRepo.Update(ctx, room)
}