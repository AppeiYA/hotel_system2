package domain

import "slices"

type ReservationStatus string

const (
	ReservationStatusPending    ReservationStatus = "pending"
	ReservationStatusConfirmed  ReservationStatus = "confirmed"
	ReservationStatusCheckedIn  ReservationStatus = "checked_in"
	ReservationStatusCheckedOut ReservationStatus = "checked_out"
	ReservationStatusCancelled  ReservationStatus = "cancelled"
	ReservationStatusNoShow     ReservationStatus = "no_show"
)

func (rs ReservationStatus) IsValid() bool {
	switch rs {
	case ReservationStatusPending, ReservationStatusConfirmed, ReservationStatusCheckedIn, ReservationStatusCheckedOut, ReservationStatusCancelled, ReservationStatusNoShow:
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
	ReservationStatusNoShow:     {}, // terminal
}

func (rs ReservationStatus) CanTransitionTo(target ReservationStatus) bool {
	allowed, ok := validReservationTransitions[rs]
	return ok && slices.Contains(allowed, target)
}