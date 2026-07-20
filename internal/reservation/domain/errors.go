package domain

import custom_errors "hotel_system2/internal/shared/errors"

var (
	ErrOverlappingReservation = custom_errors.ConflictError("overlapping reservation exists")
	ErrReservationNotConfirmed = custom_errors.BadException("reservation is not confirmed")
	ErrReservationNotFound = custom_errors.NotFoundError("reservation not found")
	ErrReservationNotCheckedIn = custom_errors.BadException("reservation is not checked in")
	ErrFolioBalanceOutstanding = custom_errors.BadException("folio balance is outstanding")
	ErrReservationAlreadyConfirmed = custom_errors.BadException("reservation is already confirmed")
	ErrInvalidCheckInWindow     = custom_errors.BadException("check-out must be after check-in")
	ErrMissingReservationFields = custom_errors.BadException("guestID and roomID are required")
	ErrInvalidTransition        = custom_errors.BadException("invalid reservation state transition")
	ErrCannotCheckInEarly       = custom_errors.BadException("cannot check in before scheduled date")
	ErrPaymentNotFound = custom_errors.NotFoundError("payment not found")
)