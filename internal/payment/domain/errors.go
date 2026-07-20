package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrPaymentFailed = custom_errors.BadException("payment failed")
	ErrPaymentAlreadyInitialized = custom_errors.ConflictError("payment already initialized")
	ErrPaymentAlreadyCompleted = custom_errors.ConflictError("payment already completed")
	ErrMissingReservationID = custom_errors.BadException("missing reservation ID")
	ErrInvalidPaymentAmount = custom_errors.BadException("invalid payment amount")
	ErrInvalidPaymentTransition = custom_errors.BadException("invalid payment state transition")
	ErrPaymentNotFound = custom_errors.NotFoundError("payment not found")
)