package ports

import (
	"context"
	"hotel_system2/internal/room/domain"
)

type RoomRepository interface {
	Create(ctx context.Context, room *domain.Room) error
	FindByID(ctx context.Context, id string) (*domain.Room, error)
	FindByIDForUpdate(ctx context.Context, id string) (*domain.Room, error)
	FindByNumber(ctx context.Context, roomNumber string) (*domain.Room, error)
	List(ctx context.Context) ([]*domain.Room, error)
	Update(ctx context.Context, room *domain.Room) error
	Delete(ctx context.Context, id string) error
}