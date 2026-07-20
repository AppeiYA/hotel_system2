package reservation_http

import (
	"fmt"
	"time"

	reservation_domain "hotel_system2/internal/reservation/domain"
	reservation_usecase "hotel_system2/internal/reservation/use_case"
	shared_domain "hotel_system2/internal/shared/domain"
)

type createReservationRequest struct {
	FirstName string                              `json:"first_name" validate:"required"`
	LastName  string                              `json:"last_name" validate:"required"`
	Email     string                              `json:"email" validate:"required,email"`
	Phone     string                              `json:"phone" validate:"required"`
	RoomID    string                              `json:"room_id" validate:"required"`
	CheckIn   reservation_usecase.FlexibleDateTime `json:"check_in" validate:"required"`
	CheckOut  reservation_usecase.FlexibleDateTime `json:"check_out" validate:"required"`
}

func (r *createReservationRequest) Validate() error {
	if !r.CheckOut.After(r.CheckIn.Time) {
		return fmt.Errorf("check_out must be after check_in")
	}
	return nil
}

func (r *createReservationRequest) toInput() (reservation_usecase.CreateReservationInput, error) {
	email, err := shared_domain.NewEmail(r.Email)
	if err != nil {
		return reservation_usecase.CreateReservationInput{}, err
	}
	return reservation_usecase.CreateReservationInput{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     email,
		Phone:     r.Phone,
		RoomID:    r.RoomID,
		CheckIn:   r.CheckIn,
		CheckOut:  r.CheckOut,
	}, nil
}

type reservationDetailsOnly struct {
	ID          string                              `json:"id"`
	GuestID     string                              `json:"guest_id"`
	RoomID      string                              `json:"room_id"`
	CheckIn     time.Time                           `json:"check_in"`
	CheckOut    time.Time                           `json:"check_out"`
	TotalAmount int64                               `json:"total_amount"`
	Status      reservation_domain.ReservationStatus `json:"status"`
	CreatedAt   time.Time                           `json:"created_at"`
}

type reservationResponse struct {
	Reservation reservationDetailsOnly `json:"reservation"`
	PaymentID   *string                `json:"payment_id,omitempty"`
}

func toReservationResponse(details *reservation_domain.ReservationDetails) reservationResponse {
	res := details.Reservation
	dr := res.DateRange()

	resp := reservationResponse{
		Reservation: reservationDetailsOnly{
			ID:          res.ID(),
			GuestID:     res.GuestID(),
			RoomID:      res.RoomID(),
			CheckIn:     dr.CheckIn,
			CheckOut:    dr.CheckOut,
			TotalAmount: res.TotalAmount().AmountMinor,
			Status:      res.Status(),
			CreatedAt:   res.CreatedAt(),
		},
	}
	if details.PaymentID != "" {
		resp.PaymentID = &details.PaymentID
	}
	return resp
}