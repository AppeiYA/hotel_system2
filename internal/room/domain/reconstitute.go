package domain

import (
	shared_domain "hotel_system2/internal/shared/domain"
	"time"
)

func ReconstituteRoom(
	id, roomNumber string,
	roomType RoomType,
	rate shared_domain.Money,
	status RoomStatus,
	createdAt time.Time,
) *Room {
	return &Room{
		id:         id,
		roomNumber: roomNumber,
		roomType:   roomType,
		rate:       rate,
		status:     status,
		createdAt:  createdAt,
	}
}