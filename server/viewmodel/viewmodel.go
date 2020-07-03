package viewmodel

import (
	"errors"
	"time"
)

// Date custom date type
type Date string

// LayoutSimpleDate date format standard value
const LayoutSimpleDate string = "2006-01-02"

func (d Date) String() string {
	return string(d)
}

// ParseDateToTime parse date to time
func ParseDateToTime(date Date) (t time.Time, err error) {
	if date.String() == "" {
		return time.Time{}, nil
	}

	t, err = time.Parse(LayoutSimpleDate, date.String())
	if err != nil {
		return t, errors.New("Invalid date format")
	}

	return t, nil
}
