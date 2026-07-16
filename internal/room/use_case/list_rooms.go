package room_usecase

import (
	"context"
	"hotel_system2/internal/room/domain"
	"hotel_system2/internal/room/ports"
)

type ListRooms struct {
	roomRepo ports.Repository
}

func NewListRooms(roomRepo ports.Repository) *ListRooms {
	return &ListRooms{roomRepo: roomRepo}
}

func (l *ListRooms) Execute(ctx context.Context) ([]domain.Room, error) {
	rooms, err := l.roomRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}