package payment_http

type initializePaymentRequest struct {
	ReservationID string `json:"reservation_id" validate:"required,uuid"`
	Email         string `json:"email" validate:"required,email"`
}

type initializePaymentResponse struct {
	Reference   string `json:"reference"`
	CheckoutURL string `json:"checkout_url"`
}

