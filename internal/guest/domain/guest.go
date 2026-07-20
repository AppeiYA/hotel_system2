package domain

import (
	"time"

	shared_domain "hotel_system2/internal/shared/domain"
)

type Guest struct {
	id        string
	firstName string
	lastName  string
	email     shared_domain.Email
	phone     string
	createdAt time.Time
}

func NewGuest(
	id, firstName, lastName string,
	email shared_domain.Email,
	phone string,
) (*Guest, error) {
	if firstName == "" || lastName == "" {
		return nil, ErrMissingGuestName
	}
	return &Guest{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		phone:     phone,
		createdAt: time.Now(),
	}, nil
}

// ---- Getters ----

func (g *Guest) ID() string                    { return g.id }
func (g *Guest) FirstName() string              { return g.firstName }
func (g *Guest) LastName() string               { return g.lastName }
func (g *Guest) Email() shared_domain.Email      { return g.email }
func (g *Guest) Phone() string                  { return g.phone }
func (g *Guest) CreatedAt() time.Time           { return g.createdAt }

// ---- Behavior ----

func (g *Guest) FullName() string {
	return g.firstName + " " + g.lastName
}

func (g *Guest) UpdateContactInfo(email shared_domain.Email, phone string) {
	g.email = email
	g.phone = phone
}