package room_postgres

import (
	shared_domain "hotel_system2/internal/shared/domain"
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

func (r roomRow) toDomain() *domain.Room {
	return domain.ReconstituteRoom(
		r.ID,
		r.RoomNumber,
		domain.RoomType(r.Type),
		shared_domain.Money{AmountMinor: r.Rate, Currency: shared_domain.DefaultCurrency},
		domain.RoomStatus(r.Status),
		r.CreatedAt,
	)
}

func roomRowFromDomain(room *domain.Room) roomRow {
	rate := room.Rate()
	return roomRow{
		ID:         room.ID(),
		RoomNumber: room.RoomNumber(),
		Type:       room.Type(),
		Rate:  rate.AmountMinor,
		Status:     room.Status(),
		CreatedAt:  room.CreatedAt(),
	}
}