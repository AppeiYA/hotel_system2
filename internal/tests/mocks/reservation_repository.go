package test_mocks

import (
	"context"
	reservation_domain "hotel_system2/internal/reservation/domain"
	"time"
)

type MockReservationRepository struct {
	CreateFn                func(ctx context.Context, reservation *reservation_domain.Reservation) error
	FindByIDFn               func(ctx context.Context, id string) (*reservation_domain.Reservation, error)
	FindByIDForUpdateFn      func(ctx context.Context, id string) (*reservation_domain.Reservation, error)
	ListFn                   func(ctx context.Context) ([]*reservation_domain.Reservation, error)
	ListByEmailFn            func(ctx context.Context, email string) ([]*reservation_domain.Reservation, error)
	UpdateFn                 func(ctx context.Context, reservation *reservation_domain.Reservation) error
	UpdateStatusFn           func(ctx context.Context, id string, status reservation_domain.ReservationStatus) error
	HasOverlapFn             func(ctx context.Context, roomID string, checkIn, checkOut time.Time) (bool, error)
	FindExpiredPendingFn     func(ctx context.Context, olderThan time.Time) ([]*reservation_domain.Reservation, error)
	FindNoShowFn             func(ctx context.Context, before time.Time) ([]*reservation_domain.Reservation, error)
	GetReservationDetailsFn  func(ctx context.Context, id string) (*reservation_domain.ReservationDetails, error)
}

func (m *MockReservationRepository) Create(ctx context.Context, reservation *reservation_domain.Reservation) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, reservation)
	}
	return nil
}

func (m *MockReservationRepository) FindByID(ctx context.Context, id string) (*reservation_domain.Reservation, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockReservationRepository) FindByIDForUpdate(ctx context.Context, id string) (*reservation_domain.Reservation, error) {
	if m.FindByIDForUpdateFn != nil {
		return m.FindByIDForUpdateFn(ctx, id)
	}
	return nil, nil
}

func (m *MockReservationRepository) List(ctx context.Context) ([]*reservation_domain.Reservation, error) {
	if m.ListFn != nil {
		return m.ListFn(ctx)
	}
	return nil, nil
}

func (m *MockReservationRepository) ListByEmail(ctx context.Context, email string) ([]*reservation_domain.Reservation, error) {
	if m.ListByEmailFn != nil {
		return m.ListByEmailFn(ctx, email)
	}
	return nil, nil
}

func (m *MockReservationRepository) Update(ctx context.Context, reservation *reservation_domain.Reservation) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, reservation)
	}
	return nil
}

func (m *MockReservationRepository) UpdateStatus(ctx context.Context, id string, status reservation_domain.ReservationStatus) error {
	if m.UpdateStatusFn != nil {
		return m.UpdateStatusFn(ctx, id, status)
	}
	return nil
}

func (m *MockReservationRepository) HasOverlap(ctx context.Context, roomID string, checkIn, checkOut time.Time) (bool, error) {
	if m.HasOverlapFn != nil {
		return m.HasOverlapFn(ctx, roomID, checkIn, checkOut)
	}
	return false, nil
}

func (m *MockReservationRepository) FindExpiredPending(ctx context.Context, olderThan time.Time) ([]*reservation_domain.Reservation, error) {
	if m.FindExpiredPendingFn != nil {
		return m.FindExpiredPendingFn(ctx, olderThan)
	}
	return nil, nil
}

func (m *MockReservationRepository) FindNoShow(ctx context.Context, before time.Time) ([]*reservation_domain.Reservation, error) {
	if m.FindNoShowFn != nil {
		return m.FindNoShowFn(ctx, before)
	}
	return nil, nil
}

func (m *MockReservationRepository) GetReservationDetails(ctx context.Context, id string) (*reservation_domain.ReservationDetails, error) {
	if m.GetReservationDetailsFn != nil {
		return m.GetReservationDetailsFn(ctx, id)
	}
	return nil, nil
}