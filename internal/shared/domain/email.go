package domain

import (
	custom_errors "hotel_system2/internal/shared/errors"
	"net/mail"
)

var (
	ErrInvalidEmail = custom_errors.BadException("invalid email format")
	ErrEmptyEmail = custom_errors.BadException("email cannot be empty")
)


type Email struct {
	value string
}

func NewEmail(raw string) (Email, error) {
	if err := Validate(raw); err != nil { // simple regex/format check
		return Email{}, err
	}
	return Email{value: raw}, nil
}

func Validate(raw string) error {
	_, err := mail.ParseAddress(raw)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}

func (e Email) String() string {
	return e.value
}