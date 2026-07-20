package room_http

import (
	"time"

	"hotel_system2/internal/room/domain"
)

type createRoomRequest struct {
	RoomNumber string `json:"room_number"`
	RoomType   string `json:"room_type"`
	RateMinor  int64  `json:"rate_minor"`
	Currency   string `json:"currency"`
}

type updateRoomRequest struct {
	RoomNumber string `json:"room_number"`
	RoomType   string `json:"room_type"`
	RateMinor  int64  `json:"rate_minor"`
	Currency   string `json:"currency"`
}

type roomResponse struct {
	ID         string    `json:"id"`
	RoomNumber string    `json:"room_number"`
	RoomType   string    `json:"room_type"`
	RateMinor  int64     `json:"rate_minor"`
	Currency   string    `json:"currency"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func toRoomResponse(r *domain.Room) roomResponse {
	rate := r.Rate()
	return roomResponse{
		ID:         r.ID(),
		RoomNumber: r.RoomNumber(),
		RoomType:   string(r.Type()),
		RateMinor:  rate.AmountMinor,
		Currency:   rate.Currency,
		Status:     string(r.Status()),
		CreatedAt:  r.CreatedAt(),
	}
}