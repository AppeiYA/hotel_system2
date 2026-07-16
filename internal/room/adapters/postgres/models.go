package room_postgres

import (
	"hotel_system2/internal/room/domain"
	"time"
)

type roomRow struct {
	ID         string          `db:"id"`
	RoomNumber string          `db:"room_number"`
	Type       domain.RoomType `db:"room_type"`
	Rate       int64           `db:"rate"`
	Status     domain.RoomStatus `db:"status"`
	CreatedAt  time.Time       `db:"created_at"`
}

func (r roomRow) toDomain() domain.Room {
	return domain.Room{
		ID:         r.ID,
		RoomNumber: r.RoomNumber,
		Type:       r.Type,
		Rate:       r.Rate,
		Status:     r.Status,
		CreatedAt:  r.CreatedAt,
	}
}

func roomRowFromDomain(room *domain.Room) roomRow {
	return roomRow{
		ID:         room.ID,
		RoomNumber: room.RoomNumber,
		Type:       room.Type,
		Rate:       room.Rate,
		Status:     room.Status,
		CreatedAt:  room.CreatedAt,
	}
}