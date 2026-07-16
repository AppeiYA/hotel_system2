package room_usecase

import (
	"context"
	"hotel_system2/internal/room/domain"
	"hotel_system2/internal/room/ports"
	custom_errors "hotel_system2/internal/shared/errors"
	"time"
)

type FindAvailableRoom struct {
	roomRepo ports.Repository
}

func NewFindAvailableRoom(roomRepo ports.Repository) *FindAvailableRoom {
	return &FindAvailableRoom{roomRepo: roomRepo}
}

func (f *FindAvailableRoom) Execute(ctx context.Context, roomType domain.RoomType, checkIn time.Time, checkOut time.Time) (*domain.Room, error) {
	if !roomType.IsValid() {
		return nil, custom_errors.BadException("room type id is required")
	}
	if checkIn.After(checkOut) || checkIn.Equal(checkOut) {
		return nil, custom_errors.BadException("check-in date must be before check-out date")
	}
	return f.roomRepo.FindAvailable(ctx, roomType, checkIn, checkOut)
}
