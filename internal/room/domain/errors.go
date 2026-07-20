package domain

import (
	custom_errors "hotel_system2/internal/shared/errors"
)

var (
	ErrRoomNotFound = custom_errors.NotFoundError("room not found")
	ErrRoomConflict  = custom_errors.ConflictError("room conflict")
	ErrRoomUnauthorized = custom_errors.NewCustomError("unauthorized access to room", 401, custom_errors.ErrUnauthorized)
	ErrRoomUnavailable = custom_errors.BadException("room is unavailable")
	ErrRoomOccupied = custom_errors.BadException("room is already occupied")
	ErrMissingRoomNumber = custom_errors.BadException("missing room number")
	ErrInvalidRoomRate = custom_errors.BadException("invalid room rate")
	ErrInvalidRoomTransition = custom_errors.BadException("invalid room state transition")
)