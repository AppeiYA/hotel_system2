package reservation_usecase

import (
	"fmt"
	"strings"
	"time"
)

var dateTimeLayouts = []string{
	time.RFC3339,          // 2026-07-07T15:00:00Z / with offset
	"2006-01-02T15:04:05", // no timezone
	"2006-01-02 15:04:05", // space-separated
	"2006-01-02T15:04",    // no seconds
	"2006-01-02",          // date only — falls back to midnight
}

type FlexibleDateTime struct {
	time.Time
}

func (d *FlexibleDateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return nil
	}

	var lastErr error
	for _, layout := range dateTimeLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			d.Time = t
			return nil
		}
		lastErr = err
	}
	return fmt.Errorf("invalid datetime %q: %w", s, lastErr)
}

func (d FlexibleDateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(time.RFC3339) + `"`), nil
}

func ComputeTotalAmount(rate int64, checkIn, checkOut time.Time) int64 {
	start := time.Date(checkIn.Year(), checkIn.Month(), checkIn.Day(), 0, 0, 0, 0, checkIn.Location())
	end := time.Date(checkOut.Year(), checkOut.Month(), checkOut.Day(), 0, 0, 0, 0, checkOut.Location())

	if end.Before(start) {
		return 0
	}

	nights := int64(end.Sub(start).Hours() / 24)

	const minimumNights = 1
	if nights < minimumNights {
		nights = minimumNights
	}

	return nights * rate
}