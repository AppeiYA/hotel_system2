package payment_http

import (
    _ "hotel_system2/internal/payment/domain"
    _ "hotel_system2/internal/reservation/domain"
    _ "hotel_system2/internal/shared/response"
)

// Initialize godoc
//
// @Summary Initialize Payment
// @Description Initializes a payment session for a reservation.
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body initializePaymentRequest true "Payment initialization request"
// @Success 201 {object} response.Response{data=initializePaymentResponse}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /payments/initialize [post]
func _Initialize() {}

// Webhook godoc
//
// @Summary Confirm Payment
// @Description Confirms a payment session for a reservation.
// @Tags Payment
// @Accept json
// @Produce json
// @Param reference path string true "Payment reference"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /payments/webhook/{reference} [post]
func _Webhook() {}