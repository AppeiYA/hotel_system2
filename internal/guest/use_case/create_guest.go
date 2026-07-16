package guest_usecase

import (
	"context"
	"hotel_system2/internal/guest/domain"
	"hotel_system2/internal/guest/ports"
	custom_errors "hotel_system2/internal/shared/errors"
)

type CreateGuest struct {
	guestRepo ports.Repository
}

func NewCreateGuest(guestRepo ports.Repository) *CreateGuest {
	return &CreateGuest{guestRepo: guestRepo}
}

func (uc *CreateGuest) Execute(
	ctx context.Context,
	input domain.Guest,
) error {
	err := uc.guestRepo.Create(ctx, &input)
	if err != nil {
		return custom_errors.InternalServerError("Error processing guest record: " + err.Error())
	}
	return nil
}