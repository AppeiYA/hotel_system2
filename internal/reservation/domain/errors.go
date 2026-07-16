package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrOverlappingReservation = custom_errors.ConflictError("overlapping reservation exists")
	ErrInvalidCheckInWindow = custom_errors.BadException("invalid check in window")
	ErrReservationNotConfirmed = custom_errors.BadException("reservation is not confirmed")
	ErrReservationNotFound = custom_errors.NotFoundError("reservation not found")
	ErrReservationNotCheckedIn = custom_errors.BadException("reservation is not checked in")
	ErrFolioBalanceOutstanding = custom_errors.BadException("folio balance is outstanding")
	ErrReservationAlreadyConfirmed = custom_errors.BadException("reservation is already confirmed")
)