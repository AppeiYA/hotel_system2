package domain

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "card"
	PaymentMethodDebitCard  PaymentMethod = "transfer"
	PaymentMethodCash       PaymentMethod = "cash"
)

func (pm PaymentMethod) IsValid() bool {
	switch pm {
	case PaymentMethodCreditCard, PaymentMethodDebitCard, PaymentMethodCash:
		return true
	default:
		return false
	}
}