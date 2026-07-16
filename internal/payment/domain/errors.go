package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrPaymentFailed = custom_errors.BadException("payment failed")
)