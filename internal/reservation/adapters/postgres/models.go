package reservation_postgres

import (
	"database/sql"
	payment_domain "hotel_system2/internal/payment/domain"
	reservation_domain "hotel_system2/internal/reservation/domain"
	"time"
)

type reservationRow struct {
	ID          string            `json:"id" db:"id"`
	GuestID     string            `json:"guest_id" db:"guest_id"`
	RoomID      string            `json:"room_id" db:"room_id"`
	CheckIn     time.Time         `json:"check_in" db:"check_in"`
	CheckOut    time.Time         `json:"check_out" db:"check_out"`
	TotalAmount int64             `json:"total_amount" db:"total_amount"`
	Status      reservation_domain.ReservationStatus `json:"status" db:"status"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
}

func (r *reservationRow) toDomain() *reservation_domain.Reservation {
	if r == nil {
		return nil
	}
	
	return &reservation_domain.Reservation{
		ID:          r.ID,
		GuestID:     r.GuestID,
		RoomID:      r.RoomID,
		CheckIn:     r.CheckIn,
		CheckOut:    r.CheckOut,
		TotalAmount: r.TotalAmount,
		Status:      r.Status,
		CreatedAt:   r.CreatedAt,
	}
}

type reservationDetailsRow struct {
	// Reservation
	ID          string                             `db:"id"`
	GuestID     string                             `db:"guest_id"`
	RoomID      string                             `db:"room_id"`
	CheckIn     time.Time                          `db:"check_in"`
	CheckOut    time.Time                          `db:"check_out"`
	TotalAmount int64                              `db:"total_amount"`
	Status      reservation_domain.ReservationStatus `db:"status"`
	CreatedAt   time.Time                          `db:"created_at"`

	// Payment (nullable because of LEFT JOIN)
	PaymentID     sql.NullString                 `db:"payment_id"`
	PaymentMethod sql.NullString                 `db:"payment_method"`
	PaymentStatus sql.NullString                 `db:"payment_status"`
	PaymentAmount sql.NullInt64                  `db:"payment_amount"`
}

func (r *reservationDetailsRow) toDomain() *reservation_domain.ReservationDetails {
	details := &reservation_domain.ReservationDetails{
		Reservation: reservation_domain.Reservation{
			ID:          r.ID,
			GuestID:     r.GuestID,
			RoomID:      r.RoomID,
			CheckIn:     r.CheckIn,
			CheckOut:    r.CheckOut,
			TotalAmount: r.TotalAmount,
			Status:      r.Status,
			CreatedAt:   r.CreatedAt,
		},
	}

	if r.PaymentID.Valid {
		details.Payment = &payment_domain.Payment{
			ID:     r.PaymentID.String,
			Method: payment_domain.PaymentMethod(r.PaymentMethod.String),
			Status: payment_domain.PaymentStatus(r.PaymentStatus.String),
			Amount: r.PaymentAmount.Int64,
		}
	}

	return details
}