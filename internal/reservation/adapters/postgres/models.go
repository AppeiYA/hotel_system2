package reservation_postgres

import (
	"time"

	reservation_domain "hotel_system2/internal/reservation/domain"
	shared_domain "hotel_system2/internal/shared/domain"
)

type reservationRow struct {
	ID          string    `db:"id"`
	GuestID     string    `db:"guest_id"`
	RoomID      string    `db:"room_id"`
	CheckIn     time.Time `db:"check_in"`
	CheckOut    time.Time `db:"check_out"`
	TotalAmount int64     `db:"total_amount"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

func (r *reservationRow) toDomain() (*reservation_domain.Reservation, error) {
	if r == nil {
		return nil, nil
	}

	dateRange, err := reservation_domain.NewDateRange(r.CheckIn, r.CheckOut)
	if err != nil {
		return nil, err
	}

	return reservation_domain.ReconstituteReservation(
		r.ID,
		r.GuestID,
		r.RoomID,
		dateRange,
		shared_domain.Money{AmountMinor: r.TotalAmount, Currency: shared_domain.DefaultCurrency},
		reservation_domain.ReservationStatus(r.Status),
		r.CreatedAt,
	), nil
}

func reservationRowFromDomain(res *reservation_domain.Reservation) reservationRow {
	dr := res.DateRange()
	return reservationRow{
		ID:          res.ID(),
		GuestID:     res.GuestID(),
		RoomID:      res.RoomID(),
		CheckIn:     dr.CheckIn,
		CheckOut:    dr.CheckOut,
		TotalAmount: res.TotalAmount().AmountMinor,
		Status:      string(res.Status()),
		CreatedAt:   res.CreatedAt(),
	}
}