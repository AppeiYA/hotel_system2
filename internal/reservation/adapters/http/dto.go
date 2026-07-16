package reservation_http

import (
	"fmt"
	payment_domain "hotel_system2/internal/payment/domain"
	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_usecase "hotel_system2/internal/reservation/use_case"
	"time"
)

type createReservationRequest struct {
	FirstName string                              `json:"first_name" validate:"required"`
	LastName  string                              `json:"last_name" validate:"required"`
	Email     string                              `json:"email" validate:"required,email"`
	Phone     string                              `json:"phone" validate:"required"`
	RoomID    string                              `json:"room_id" validate:"required"`
	CheckIn   reservation_usecase.FlexibleDateTime `json:"check_in" validate:"required"`
	CheckOut  reservation_usecase.FlexibleDateTime              `json:"check_out" validate:"required"`
}

func (r *createReservationRequest) Validate() error {
	if !r.CheckOut.After(r.CheckIn.Time) {
		return fmt.Errorf("check_out must be after check_in")
	}
	return nil
}

type reservationDetailsOnly struct {
	ID          string `json:"id" db:"id"`
	GuestID     string `json:"guest_id" db:"guest_id"`
	RoomID      string `json:"room_id" db:"room_id"`
	CheckIn     time.Time `json:"check_in" db:"check_in"`
	CheckOut    time.Time `json:"check_out" db:"check_out"`
	TotalAmount int64 `json:"total_amount" db:"total_amount"`
	Status      reservation_domain.ReservationStatus `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type paymentDetailsOnly struct {
	ID            string        `json:"id" db:"id"`
	ReservationID string        `json:"reservation_id" db:"reservation_id"`
	Reference     string        `json:"reference" db:"reference"`
	Amount        int64         `json:"amount" db:"amount"`
	Method        payment_domain.PaymentMethod `json:"method" db:"method"`
	Status        payment_domain.PaymentStatus `json:"status" db:"status"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
}

type reservationResponse struct {
	Reservation reservationDetailsOnly `json:"reservation"`
	Payment     *paymentDetailsOnly        `json:"payment,omitempty"`
}
