package sample

import (
	"encoding/json"
	"errors"
)

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
	case "Unknown", "unknown", "UNKNOWN", "":
		return UnknownSex, nil
	default:
		return UnknownSex, errors.New("Invalid Sex")
	}
}

// String returns sex string.
func (s *Sex) String() string {
	if s == nil {
		return ""
	}
	switch *s {
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
