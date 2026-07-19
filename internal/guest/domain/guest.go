package domain

import "time"

type Guest struct {
	ID           string   
	FirstName    string  
	LastName     string 
	Email        Email 
	Phone        string 
	CreatedAt    time.Time
}