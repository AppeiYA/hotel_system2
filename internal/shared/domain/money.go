// shared/domain/money.go
package domain

import "errors"

const DefaultCurrency = "NGN"

var ErrUnsupportedCurrency = errors.New("only NGN is supported")

type Money struct {
	AmountMinor int64
	Currency    string
}

func NewMoney(amountMinor int64) (Money, error) {
	if amountMinor < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	return Money{AmountMinor: amountMinor, Currency: DefaultCurrency}, nil
}

func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, ErrUnsupportedCurrency
	}
	return Money{AmountMinor: m.AmountMinor + other.AmountMinor, Currency: m.Currency}, nil
}

func (m Money) MultiplyByNights(nights int) (Money, error) {
	if nights < 0 {
		return Money{}, errors.New("nights cannot be negative")
	}
	return Money{AmountMinor: m.AmountMinor * int64(nights), Currency: m.Currency}, nil
}