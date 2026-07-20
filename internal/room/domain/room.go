package domain

import (
	"time"
	shared_domain "hotel_system2/internal/shared/domain"
)

type Room struct {
	id         string
	roomNumber string
	roomType   RoomType
	rate       shared_domain.Money
	status     RoomStatus
	createdAt  time.Time
}

func NewRoom(
	id, roomNumber string,
	roomType RoomType,
	rate shared_domain.Money,
) (*Room, error) {
	if roomNumber == "" {
		return nil, ErrMissingRoomNumber
	}
	if rate.AmountMinor <= 0 {
		return nil, ErrInvalidRoomRate
	}
	return &Room{
		id:         id,
		roomNumber: roomNumber,
		roomType:   roomType,
		rate:       rate,
		status:     RoomStatusAvailable,
		createdAt:  time.Now(),
	}, nil
}

// ---- Getters ----

func (r *Room) ID() string                 { return r.id }
func (r *Room) RoomNumber() string          { return r.roomNumber }
func (r *Room) Type() RoomType              { return r.roomType }
func (r *Room) Rate() shared_domain.Money    { return r.rate }
func (r *Room) Status() RoomStatus          { return r.status }
func (r *Room) CreatedAt() time.Time        { return r.createdAt }

// ---- State transitions ----

func (r *Room) Occupy() error {
	if r.status != RoomStatusAvailable {
		return ErrInvalidRoomTransition
	}
	r.status = RoomStatusOccupied
	return nil
}

// MarkForCleaning is the checkout path — a guest leaves,
// housekeeping needs to turn the room over before it's bookable again.
func (r *Room) MarkForCleaning() error {
	if r.status != RoomStatusOccupied {
		return ErrInvalidRoomTransition
	}
	r.status = RoomStatusCleaning
	return nil
}

func (r *Room) MarkAvailable() error {
	if r.status != RoomStatusCleaning && r.status != RoomStatusMaintenance {
		return ErrInvalidRoomTransition
	}
	r.status = RoomStatusAvailable
	return nil
}

// SendToMaintenance only makes sense from available or cleaning —
// pulling an occupied room into maintenance mid-stay is an
// operational exception, not a normal transition, so it's
// deliberately not allowed here.
func (r *Room) SendToMaintenance() error {
	if r.status != RoomStatusAvailable && r.status != RoomStatusCleaning {
		return ErrInvalidRoomTransition
	}
	r.status = RoomStatusMaintenance
	return nil
}