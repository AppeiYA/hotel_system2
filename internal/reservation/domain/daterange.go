package domain

import "time"

type DateRange struct {
	CheckIn  time.Time
	CheckOut time.Time
}

func NewDateRange(checkIn, checkOut time.Time) (DateRange, error) {
	if !checkOut.After(checkIn) {
		return DateRange{}, ErrInvalidCheckInWindow
	}
	return DateRange{CheckIn: checkIn, CheckOut: checkOut}, nil
}

func (d DateRange) Overlaps(other DateRange) bool {
	return d.CheckIn.Before(other.CheckOut) && other.CheckIn.Before(d.CheckOut)
}

func (d DateRange) Nights() int {
	return int(d.CheckOut.Sub(d.CheckIn).Hours() / 24)
}