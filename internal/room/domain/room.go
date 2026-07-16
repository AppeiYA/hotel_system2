package domain

import "time"

type Room struct {
	ID         string     
	RoomNumber string     
	Type       RoomType   
	Rate       int64    
	Status     RoomStatus 
	CreatedAt  time.Time
}