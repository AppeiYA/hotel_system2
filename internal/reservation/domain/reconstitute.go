package domain

import (
	"time"

	shared_domain "hotel_system2/internal/shared/domain"
)

func ReconstituteReservation(
	id, guestID, roomID string,
	dateRange DateRange,
	totalAmount shared_domain.Money,
	status ReservationStatus,
	createdAt time.Time,
) *Reservation {
	return &Reservation{
		id:          id,
		guestID:     guestID,
		roomID:      roomID,
		dateRange:   dateRange,
		totalAmount: totalAmount,
		status:      status,
		createdAt:   createdAt,
	}
}