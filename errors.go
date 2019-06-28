package sheetdb

import "fmt"

type NotFoundError struct {
	Model string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Model '%s' not found", e.Model)
}

type EmptyStringError struct {
	FieldName string
}

func (e *EmptyStringError) Error() string {
	return fmt.Sprintf("Value '%s' is empty", e.FieldName)
}

type DuplicationError struct {
	FieldName string
}

func (e *DuplicationError) Error() string {
	return fmt.Sprintf("Value '%s' is duplicated", e.FieldName)
}

type InvalidValueError struct {
	FieldName string
	Err       error
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("Value '%s' is invalid: %s", e.FieldName, e.Err)
}
