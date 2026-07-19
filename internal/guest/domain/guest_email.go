package domain

import (
	custom_errors "hotel_system2/internal/shared/errors"
	"net/mail"
	"strings"
)

type Email string

var (
	ErrInvalidEmail = custom_errors.BadException("invalid email format")
	ErrEmptyEmail = custom_errors.BadException("email cannot be empty")
)

func (e Email) Validate() error {
	s := strings.TrimSpace(string(e))
	if s == "" {
		return ErrEmptyEmail
	}
	if len(s) > 254 {
		return ErrInvalidEmail
	}

	_, err := mail.ParseAddress(s)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}