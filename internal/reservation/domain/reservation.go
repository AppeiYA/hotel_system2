package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
	"time"
)

type Reservation struct {
	id          string
	guestID     string
	roomID      string
	dateRange   DateRange
	totalAmount shared_domain.Money
	status      ReservationStatus
	createdAt   time.Time
}

func NewReservation(
	id, guestID, roomID string,
	dr DateRange,
	amount shared_domain.Money,
) (*Reservation, error) {
	if guestID == "" || roomID == "" {
		return nil, ErrMissingReservationFields
	}
	return &Reservation{
		id:          id,
		guestID:     guestID,
		roomID:      roomID,
		dateRange:   dr,
		totalAmount: amount,
		status:      ReservationStatusPending,
		createdAt:   time.Now(),
	}, nil
}

// ---- Getters (read-only access from outside the package) ----
func (r *Reservation) ID() string                        { return r.id }
func (r *Reservation) GuestID() string                   { return r.guestID }
func (r *Reservation) RoomID() string                    { return r.roomID }
func (r *Reservation) DateRange() DateRange               { return r.dateRange }
func (r *Reservation) TotalAmount() shared_domain.Money    { return r.totalAmount }
func (r *Reservation) Status() ReservationStatus          { return r.status }
func (r *Reservation) CreatedAt() time.Time               { return r.createdAt }

// ---- State transitions----
// func (r *Reservation) Confirm() error {
// 	if r.status != ReservationStatusPending {
// 		return ErrInvalidTransition
// 	}
// 	r.status = ReservationStatusConfirmed
// 	return nil
// }

// func (r *Reservation) CheckIn(now time.Time) error {
// 	if r.status != ReservationStatusConfirmed {
// 		return ErrInvalidTransition
// 	}
// 	if now.Before(r.dateRange.CheckIn) {
// 		return ErrCannotCheckInEarly
// 	}
// 	r.status = ReservationStatusCheckedIn
// 	return nil
// }

// func (r *Reservation) CheckOut() error {
// 	if r.status != ReservationStatusCheckedIn {
// 		return ErrInvalidTransition
// 	}
// 	r.status = ReservationStatusCheckedOut
// 	return nil
// }

// func (r *Reservation) Cancel() error {
// 	if r.status == ReservationStatusCheckedIn || r.status == ReservationStatusCheckedOut {
// 		return ErrInvalidTransition
// 	}
// 	r.status = ReservationStatusCancelled
// 	return nil
// }

func (r *Reservation) transitionTo(target ReservationStatus) error {
	if !r.status.CanTransitionTo(target) {
		return ErrInvalidTransition
	}
	r.status = target
	return nil
}

func (r *Reservation) Confirm() error {
	return r.transitionTo(ReservationStatusConfirmed)
}

func (r *Reservation) CheckIn(now time.Time) error {
	if now.Before(r.dateRange.CheckIn) {
		return ErrCannotCheckInEarly
	}
	return r.transitionTo(ReservationStatusCheckedIn)
}

func (r *Reservation) CheckOut() error {
	return r.transitionTo(ReservationStatusCheckedOut)
}

func (r *Reservation) Cancel() error {
	return r.transitionTo(ReservationStatusCancelled)
}

func (r *Reservation) MarkNoShow() error {
	return r.transitionTo(ReservationStatusNoShow)
}

func (r *Reservation) MarkPending() error {
	return r.transitionTo(ReservationStatusPending)
}

type ReservationDetails struct {
	Reservation Reservation
	PaymentID   string
}