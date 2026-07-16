package domain

import (
	"hotel_system2/internal/payment/domain"
	"time"
)

type Reservation struct {
	ID          string
	GuestID     string
	RoomID      string
	CheckIn     time.Time
	CheckOut    time.Time
	TotalAmount int64
	Status      ReservationStatus
	CreatedAt   time.Time
}

type ReservationDetails struct {
    Reservation Reservation
    Payment     *domain.Payment
}