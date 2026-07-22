package domain

type ReferenceType string

const (
	ReferenceTypeReservation ReferenceType = "reservation"
	ReferenceTypePayment     ReferenceType = "payment"
)