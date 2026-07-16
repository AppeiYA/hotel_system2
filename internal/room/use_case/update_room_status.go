package room_usecase

import (
	"context"
	"hotel_system2/internal/room/domain"
	"hotel_system2/internal/room/ports"
)

type UpdateRoomStatus struct {
	roomRepo ports.Repository
}

func NewUpdateRoomStatus(roomRepo ports.Repository) *UpdateRoomStatus {
	return &UpdateRoomStatus{roomRepo: roomRepo}
}

func (uc *UpdateRoomStatus) Execute(
	ctx context.Context,
	roomID string,
	status domain.RoomStatus,
) error {

	room, err := uc.roomRepo.FindByID(ctx, roomID)
	if err != nil {
		return err
	}

	room.Status = status

	return uc.roomRepo.Update(ctx, room)
}

