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
func NewDate(date string) (Date, error) {
	d, err := time.Parse(dateFmt, date)
	if err != nil {
		return Date{}, err
	}
	return Date{d}, nil
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
	return Datetime{d}, nil
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
