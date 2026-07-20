package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
	"time"
)

func ReconstituteGuest(
	id, firstName, lastName string,
	email shared_domain.Email,
	phone string,
	createdAt time.Time,
) *Guest {
	return &Guest{
		id:         id,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		phone:      phone,
		createdAt:  createdAt,
	}
}