package payment_http

//@Summary Initialize Payment
//@Description Initializes a payment session for a reservation.
//@Tags Payment
//@Accept json
//@Produce json
//@Param request body initializePaymentRequest true "Payment initialization request"
//@Success 201 {object} response.Response{data=initializePaymentResponse}
func _Initialize(){}

//@Summary Confirm Payment
//@Description Confirms a payment session for a reservation.
//@Tags Payment
//@Accept json
//@Produce json
//@Param reference path string true "Payment reference"
//@Success 200 {object} response.Response{data=confirmPaymentResponse}
func _Complete(){}