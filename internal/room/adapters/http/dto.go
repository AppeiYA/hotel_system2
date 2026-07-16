package room_http

import (
	"hotel_system2/internal/room/domain"
	"time"
)

type createRoomRequest struct {
	RoomNumber string `json:"room_number"`
	RoomType   string `json:"room_type"`
	Rate       int64  `json:"rate"`
}

type updateRoomRequest struct {
	RoomNumber string `json:"room_number"`
	RoomType   string `json:"room_type"`
	Rate       int64  `json:"rate"`
	Status     string `json:"status"`
}

type updateStatusRequest struct {
	Status string `json:"status"`
}

type roomResponse struct {
	ID         string    `json:"id"`
	RoomNumber string    `json:"room_number"`
	RoomType   string    `json:"room_type"`
	Rate       int64     `json:"rate"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func toRoomResponse(r *domain.Room) roomResponse {
	return roomResponse{
		ID:         r.ID,
		RoomNumber: r.RoomNumber,
		RoomType:   string(r.Type),
		Rate:       r.Rate,
		Status:     string(r.Status),
		CreatedAt:  r.CreatedAt,
	}
}
