package room_usecase

import (
	"context"
	"hotel_system2/internal/room/domain"
	"hotel_system2/internal/room/ports"
	custom_errors "hotel_system2/internal/shared/errors"
)

type CreateRoom struct {
	roomRepo ports.Repository
}

func NewCreateRoom(roomRepo ports.Repository) *CreateRoom {
	return &CreateRoom{roomRepo: roomRepo}
}

func (c *CreateRoom) Execute(ctx context.Context, room *domain.Room) error {
	if room.RoomNumber == "" {
		return custom_errors.BadException("room number is required")
	}
	if !room.Type.IsValid() {
		return custom_errors.BadException("invalid room type")
	}
	if room.Rate <= 0 {
		return custom_errors.BadException("rate must be greater than 0")
	}

	return c.roomRepo.Create(ctx, room)
}