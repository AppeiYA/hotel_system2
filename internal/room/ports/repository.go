package ports

import (
	"context"
	"hotel_system2/internal/room/domain"
	"time"
)

type Repository interface {
	Create(ctx context.Context, room *domain.Room) error
	FindByID(ctx context.Context, id string) (*domain.Room, error)
	FindByNumber(ctx context.Context, roomNumber string) (*domain.Room, error)
	List(ctx context.Context) ([]domain.Room, error)
	Update(ctx context.Context, room *domain.Room) error
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, status domain.RoomStatus) error
	FindAvailable(
		ctx context.Context,
		roomType domain.RoomType,
		checkIn time.Time,
		checkOut time.Time,
	) (*domain.Room, error)
	FindByIDForUpdate(ctx context.Context, id string) (*domain.Room, error)
}