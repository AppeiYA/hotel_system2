package domain

import "slices"

type ReservationStatus string

const (
	ReservationStatusPending    ReservationStatus = "pending"
	ReservationStatusConfirmed  ReservationStatus = "confirmed"
	ReservationStatusCheckedIn  ReservationStatus = "checked_in"
	ReservationStatusCheckedOut ReservationStatus = "checked_out"
	ReservationStatusCancelled  ReservationStatus = "cancelled"
)

func (rs ReservationStatus) IsValid() bool {
	switch rs {
	case ReservationStatusPending, ReservationStatusConfirmed, ReservationStatusCheckedIn, ReservationStatusCheckedOut, ReservationStatusCancelled:
		return true
	default:
		return false
	}
}

var validReservationTransitions = map[ReservationStatus][]ReservationStatus{
	ReservationStatusPending:    {ReservationStatusConfirmed, ReservationStatusCancelled},
	ReservationStatusConfirmed:  {ReservationStatusCheckedIn, ReservationStatusCancelled},
	ReservationStatusCheckedIn:  {ReservationStatusCheckedOut},
	ReservationStatusCheckedOut: {}, // terminal
	ReservationStatusCancelled:  {}, // terminal
}

func (rs ReservationStatus) CanTransitionTo(target ReservationStatus) bool {
	allowed, ok := validReservationTransitions[rs]
	if !ok {
		return false
	}
	if slices.Contains(allowed, target) {
		return true
	}

	return false
}