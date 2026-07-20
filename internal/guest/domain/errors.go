package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrMissingGuestName = custom_errors.BadException("missing guest name")
	ErrGuestNotFound = custom_errors.NotFoundError("guest not found")

)