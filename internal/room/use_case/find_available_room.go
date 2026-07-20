package room_usecase

import (
	"context"
	"time"

	"hotel_system2/internal/room/domain"
	roomports "hotel_system2/internal/room/ports"
	reservationports "hotel_system2/internal/reservation/ports"
	custom_errors "hotel_system2/internal/shared/errors"
)

type FindAvailableRoom struct {
	roomRepo        roomports.RoomRepository
	reservationRepo reservationports.ReservationRepository
}

func NewFindAvailableRoom(
	roomRepo roomports.RoomRepository,
	reservationRepo reservationports.ReservationRepository,
) *FindAvailableRoom {
	return &FindAvailableRoom{roomRepo: roomRepo, reservationRepo: reservationRepo}
}

func (f *FindAvailableRoom) Execute(
	ctx context.Context,
	roomType domain.RoomType,
	checkIn time.Time,
	checkOut time.Time,
) ([]*domain.Room, error) {
	if !roomType.IsValid() {
		return nil, custom_errors.BadException("room type id is required")
	}
	if checkIn.After(checkOut) || checkIn.Equal(checkOut) {
		return nil, custom_errors.BadException("check-in date must be before check-out date")
	}

	allRooms, err := f.roomRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var available []*domain.Room
	for _, room := range allRooms {
		if room.Type() != roomType {
			continue
		}
		overlaps, err := f.reservationRepo.HasOverlap(ctx, room.ID(), checkIn, checkOut)
		if err != nil {
			return nil, err
		}
		if !overlaps {
			available = append(available, room)
		}
	}
	return available, nil
}