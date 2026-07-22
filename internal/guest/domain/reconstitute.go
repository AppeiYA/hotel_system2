package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
	"time"
)

func ReconstituteGuest(
	id, firstName, lastName string,
	email string,
	phone string,
	createdAt time.Time,
) (*Guest, error){
	validEmail, err := shared_domain.NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &Guest{
		id:         id,
		firstName:  firstName,
		lastName:   lastName,
		email:      validEmail,
		phone:      phone,
		createdAt:  createdAt,
	}, nil
}