package reservation_http

import (
	_ "hotel_system2/internal/payment/domain"
	_ "hotel_system2/internal/reservation/domain"
	_ "hotel_system2/internal/shared/response"
	"time"
)

type CreateReservationRequest struct {
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	RoomID    string    `json:"room_id" validate:"required"`
	CheckIn   time.Time `json:"check_in" validate:"required"`
	CheckOut  time.Time `json:"check_out" validate:"required"`
}

// CreateReservation godoc
//
//	@Summary		Create a new reservation
//	@Description	Create a reservation for a guest
//	@Tags			Reservations
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateReservationRequest	true	"Reservation details"
//	@Success		201		{object}	response.Response{data=reservationResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		422		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/reservations [post]
func _CreateReservation() {}

// GetReservationsByEmail godoc
//
//	@Summary		List reservations by guest email
//	@Description	Retrieve all reservations belonging to a guest email
//	@Tags			Reservations
//	@Produce		json
//	@Param			email	path		string	true	"Guest email"
//	@Success		200		{object}	response.Response{data=[]reservationResponse}
//	@Failure		400		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/reservations/{email} [get]
func _GetReservationsByEmail() {}

// ListReservations godoc
//
//	@Summary		List all reservations
//	@Description	Retrieve all reservations
//	@Tags			Reservations
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]reservationResponse}
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/reservations [get]
func _ListReservations() {}

// CheckInReservation godoc
//
//	@Summary		Check in guest
//	@Description	Check a guest into a reservation
//	@Tags			Reservations
//	@Produce		json
//	@Param			id	path		string	true	"Reservation ID (UUID)"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Failure		409	{object}	response.ErrorResponse
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/reservations/{id}/check-in [post]
func _CheckInReservation() {}

// CheckOutReservation godoc
//
//	@Summary		Check out guest
//	@Description	Check out a guest after ensuring the folio balance is settled
//	@Tags			Reservations
//	@Produce		json
//	@Param			id	path		string	true	"Reservation ID (UUID)"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Failure		409	{object}	response.ErrorResponse
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/reservations/{id}/check-out [post]
func _CheckOutReservation() {}