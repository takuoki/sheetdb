package sheetdb

import "fmt"

// InvalidFormatError means that the format is invalid.
type InvalidFormatError struct {
	Err error
}

func (e *InvalidFormatError) Error() string {
	return fmt.Sprintf("Format is invalid: %v", e.Err)
}

// NotFoundError means that model is not found.
type NotFoundError struct {
	Model string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Model '%s' not found", e.Model)
}

// EmptyStringError means that string field that does not allow empty is empty.
type EmptyStringError struct {
	FieldName string
}

func (e *EmptyStringError) Error() string {
	return fmt.Sprintf("Value '%s' is empty", e.FieldName)
}

// DuplicationError means that field value that should be unique is duplicated.
type DuplicationError struct {
	FieldName string
}

func (e *DuplicationError) Error() string {
	return fmt.Sprintf("Value '%s' is duplicated", e.FieldName)
}

// InvalidValueError means that field value is invalid.
type InvalidValueError struct {
	FieldName string
	Err       error
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("Value '%s' is invalid: %v", e.FieldName, e.Err)
}
