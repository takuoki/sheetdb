package sheetdb

import (
	"encoding/json"
	"time"
)

const dateFmt = "2006-01-02" // YYYY-MM-DD

// Date is date (YYYY-MM-DD).
type Date struct {
	time.Time
}

// NewDate returns new Date.
// The argument "date" is parsed on local location.
func NewDate(date string) (Date, error) {
	d, err := time.Parse(dateFmt, date)
	if err != nil {
		return Date{}, err
	}
	return Date{d.UTC()}, nil
}

// Today returns new Date that represents today.
func Today() Date {
	d, _ := NewDate(time.Now().Format(dateFmt))
	return d
}

// String returns date string.
func (d *Date) String() string {
	if d == nil {
		return ""
	}
	return d.Format(dateFmt)
}

// MarshalJSON marshals date value.
func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

const datetimeFmt = time.RFC3339 // YYYY-MM-DDTHH:mm:ssZ0700

// Datetime is datetime (YYYY-MM-DDTHH:mm:ssZ0700).
type Datetime struct {
	time.Time
}

// NewDatetime returns new Datetime.
func NewDatetime(datetime string) (Datetime, error) {
	d, err := time.Parse(datetimeFmt, datetime)
	if err != nil {
		return Datetime{}, err
	}
	return Datetime{d.UTC()}, nil
}

// Now returns new Datetime that represents now.
func Now() Datetime {
	d, _ := NewDatetime(time.Now().Format(datetimeFmt))
	return d
}

// String returns datetime string.
func (d *Datetime) String() string {
	if d == nil {
		return ""
	}
	return d.Format(datetimeFmt)
}

// MarshalJSON marshals datetime value.
func (d *Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
