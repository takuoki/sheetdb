package sample

import (
	"encoding/json"
)

//go:generate stringer -type=Sex

// Sex is user's sex.
type Sex int

// Sex Enumuration.
const (
	UnknownSex Sex = iota
	Male
	Female
)

// NewSex returns new Sex.
func NewSex(sex string) (Sex, error) {
	switch sex {
	case "Male", "male", "MALE":
		return Male, nil
	case "Female", "female", "FEMALE":
		return Female, nil
	default:
		return UnknownSex, nil
	}
}

// String returns sex string.
func (s Sex) String() string {
	switch s {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

// MarshalJSON marshals sex value.
func (s Sex) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
