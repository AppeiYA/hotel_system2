package domain

import "time"

type Guest struct {
	ID           string   
	FirstName    string  
	LastName     string 
	Email        string 
	Phone        string 
	CreatedAt    time.Time
}