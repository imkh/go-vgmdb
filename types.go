package vgmdb

import (
	"strconv"
	"time"
)

const (
	dtwsLayout = "2006-01-02T15:04"
)

// DateTimeWithoutSeconds represents a date (year, month, day) and time (hour, minute).
type DateTimeWithoutSeconds struct {
	time.Time
}

func (d *DateTimeWithoutSeconds) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(d.Time.Format(dtwsLayout))), nil
}

func (d *DateTimeWithoutSeconds) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		s = string(data)
	}
	if s == "null" {
		d.Time = time.Time{}
		return nil
	}
	d.Time, err = time.Parse(dtwsLayout, s)
	return err
}
