package test_mocks

import (
	"context"
	guest_domain "hotel_system2/internal/guest/domain"	
)

type MockGuestRepository struct {
	FindOrCreateFn func(ctx context.Context, guest *guest_domain.Guest) error
	CreateFn func(ctx context.Context, guest *guest_domain.Guest) error
	FindByEmailFn func(ctx context.Context, email string) (*guest_domain.Guest, error)
	ExistsByEmailFn func(ctx context.Context, email string) (bool, error)
	FindByIDFn func(ctx context.Context, id string) (*guest_domain.Guest, error)
}

func (m *MockGuestRepository) FindOrCreate(ctx context.Context, guest *guest_domain.Guest) error {
	if m.FindOrCreateFn != nil {
		return m.FindOrCreateFn(ctx, guest)
	}
	return nil

}

func (m *MockGuestRepository) Create(ctx context.Context, guest *guest_domain.Guest) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, guest)
	}
	return nil
}

func (m *MockGuestRepository) FindByEmail(ctx context.Context, email string) (*guest_domain.Guest, error) {
	if m.FindByEmailFn != nil {
		return m.FindByEmailFn(ctx, email)
	}
	return nil, nil
}

func (m *MockGuestRepository) FindByID(ctx context.Context, id string) (*guest_domain.Guest, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	return nil, nil
}


func (m *MockGuestRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	if m.ExistsByEmailFn != nil {
		return m.ExistsByEmailFn(ctx, email)
	}
	return false, nil
}