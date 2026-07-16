package reservation_http

import (
	_ "hotel_system2/internal/payment/domain"
	_ "hotel_system2/internal/reservation/domain"
	_ "hotel_system2/internal/shared/response"
	"time"
)

type CreateReservationRequest struct {
	FirstName string                              `json:"first_name" validate:"required"`
	LastName  string                              `json:"last_name" validate:"required"`
	Email     string                              `json:"email" validate:"required,email"`
	Phone     string                              `json:"phone" validate:"required"`
	RoomID    string                              `json:"room_id" validate:"required"`
	CheckIn   time.Time `json:"check_in" validate:"required"`
	CheckOut  time.Time              `json:"check_out" validate:"required"`
}

// CreateReservation godoc
//
// @Summary Create a new reservation
// @Description Create a new reservation with the provided details
// @Tags Reservations
// @Accept json
// @Produce json
// @Param request body CreateReservationRequest true "Reservation details"
// @Success 201 {object} response.Response{data=reservationResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations [post]
func _CreateReservation() {}

// GetReservation godoc
//
// @Summary Get reservation details
// @Description Get reservation details by reservation ID
// @Tags Reservations
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} reservationResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/{id} [get]
func _GetReservation() {}

// ListReservations godoc
//
// @Summary List all reservations
// @Description Get a list of all reservations in the system
// @Tags Reservations
// @Produce json
// @Success 200 {object} response.Response{data=[]reservationResponse}
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations [get]
func _ListReservations() {}

// CheckInReservation godoc
//
// @Summary Check in a reservation
// @Description Mark a confirmed reservation as checked-in and update the room status to occupied
// @Tags Reservations
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorResponse "Invalid reservation ID"
// @Failure 404 {object} response.ErrorResponse "Reservation not found"
// @Failure 409 {object} response.ErrorResponse "Conflict: Reservation not in a check-in-able state"
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/{id}/check-in [post]
func _CheckInReservation() {}

// CheckOutReservation godoc
//
// @Summary Check out a reservation
// @Description Mark a checked-in reservation as checked-out, provided the folio is settled
// @Tags Reservations
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorResponse "Invalid reservation ID"
// @Failure 404 {object} response.ErrorResponse "Reservation not found"
// @Failure 409 {object} response.ErrorResponse "Conflict: Reservation not in a check-out-able state or folio balance is outstanding"
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/{id}/check-out [post]
func _CheckOutReservation() {}