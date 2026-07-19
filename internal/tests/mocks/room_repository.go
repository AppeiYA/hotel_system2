package test_mocks

import (
	"context"
	room_domain "hotel_system2/internal/room/domain"
	"time"
)

type MockRoomRepository struct {
	CreateFn            func(ctx context.Context, room *room_domain.Room) error
	FindByIDFn          func(ctx context.Context, id string) (*room_domain.Room, error)
	FindByNumberFn      func(ctx context.Context, roomNumber string) (*room_domain.Room, error)
	ListFn              func(ctx context.Context) ([]room_domain.Room, error)
	UpdateFn            func(ctx context.Context, room *room_domain.Room) error
	DeleteFn            func(ctx context.Context, id string) error
	UpdateStatusFn      func(ctx context.Context, id string, status room_domain.RoomStatus) error
	FindAvailableFn     func(ctx context.Context, roomType room_domain.RoomType, checkIn, checkOut time.Time) (*room_domain.Room, error)
	FindByIDForUpdateFn func(ctx context.Context, id string) (*room_domain.Room, error)
}

func (m *MockRoomRepository) Create(ctx context.Context, room *room_domain.Room) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, room)
	}
	return nil
}

func (m *MockRoomRepository) FindByID(ctx context.Context, id string) (*room_domain.Room, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockRoomRepository) FindByNumber(ctx context.Context, roomNumber string) (*room_domain.Room, error) {
	if m.FindByNumberFn != nil {
		return m.FindByNumberFn(ctx, roomNumber)
	}
	return nil, nil
}

func (m *MockRoomRepository) List(ctx context.Context) ([]room_domain.Room, error) {
	if m.ListFn != nil {
		return m.ListFn(ctx)
	}
	return nil, nil
}

func (m *MockRoomRepository) Update(ctx context.Context, room *room_domain.Room) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, room)
	}
	return nil
}

func (m *MockRoomRepository) Delete(ctx context.Context, id string) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func (m *MockRoomRepository) UpdateStatus(ctx context.Context, id string, status room_domain.RoomStatus) error {
	if m.UpdateStatusFn != nil {
		return m.UpdateStatusFn(ctx, id, status)
	}
	return nil
}

func (m *MockRoomRepository) FindAvailable(ctx context.Context, roomType room_domain.RoomType, checkIn, checkOut time.Time) (*room_domain.Room, error) {
	if m.FindAvailableFn != nil {
		return m.FindAvailableFn(ctx, roomType, checkIn, checkOut)
	}
	return nil, nil
}

func (m *MockRoomRepository) FindByIDForUpdate(ctx context.Context, id string) (*room_domain.Room, error) {
	if m.FindByIDForUpdateFn != nil {
		return m.FindByIDForUpdateFn(ctx, id)
	}
	return nil, nil
}