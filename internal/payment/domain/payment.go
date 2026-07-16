package domain

import "time"

type Payment struct {
	ID            string      
	ReservationID string       
	Reference     string       
	Amount        int64        
	Method        PaymentMethod 
	Status        PaymentStatus 
	CreatedAt     time.Time   
}