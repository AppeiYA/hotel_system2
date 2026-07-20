package guest_usecase

import (
	"context"

	"hotel_system2/internal/guest/domain"
	"hotel_system2/internal/guest/ports"
	shared_domain "hotel_system2/internal/shared/domain"
)

type CreateGuestInput struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type CreateGuest struct {
	guestRepo ports.GuestRepository
}

func NewCreateGuest(guestRepo ports.GuestRepository) *CreateGuest {
	return &CreateGuest{guestRepo: guestRepo}
}

func (uc *CreateGuest) Execute(ctx context.Context, input CreateGuestInput) (*domain.Guest, error) {
	email, err := shared_domain.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	guest, err := domain.NewGuest(input.ID, input.FirstName, input.LastName, email, input.Phone)
	if err != nil {
		return nil, err
	}

	if err := uc.guestRepo.Create(ctx, guest); err != nil {
		return nil, err
	}

	return guest, nil
}