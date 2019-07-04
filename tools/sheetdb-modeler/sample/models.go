package sample

import (
	"github.com/takuoki/sheetdb"
)

var dbClient *sheetdb.Client

//go:generate sheetdb-modeler -type=User -children=Foo,FooChild,Bar -initial=10001

// User is a struct of user.
type User struct {
	UserID   int           `json:"user_id" db:"primarykey"`
	Name     string        `json:"name"`
	Email    string        `json:"email" db:"unique"`
	Sex      Sex           `json:"sex"`
	Birthday *sheetdb.Date `json:"birthday"`
}

//go:generate sheetdb-modeler -type=Foo -parent=User -children=FooChild

// Foo is a struct of foo which is a child of user.
type Foo struct {
	UserID int     `json:"user_id" db:"primarykey"`
	FooID  int     `json:"foo_id" db:"primarykey"`
	Value  float32 `json:"value"`
	Note   string  `json:"note" db:"allowempty"`
}

//go:generate sheetdb-modeler -type=FooChild -parent=Foo

// FooChild is a struct of foo child.
type FooChild struct {
	UserID  int     `json:"user_id" db:"primarykey"`
	FooID   int     `json:"foo_id" db:"primarykey"`
	ChildID int     `json:"child_id" db:"primarykey"`
	Value   float32 `json:"value"`
}

//go:generate sheetdb-modeler -type=Bar -parent=User

// Bar is a struct of bar which is a child of user.
type Bar struct {
	UserID   int              `json:"user_id" db:"primarykey"`
	Datetime sheetdb.Datetime `json:"datetime" db:"primarykey"`
	Value    float32          `json:"value"`
	Note     string           `json:"note" db:"allowempty"`
}

//go:generate sheetdb-modeler -type=TypeTest

// TypeTest is a struct for type test.
type TypeTest struct {
	ID             int               `json:"id" db:"primarykey"`
	StringValue    string            `json:"string_value"`
	BoolValue      bool              `json:"bool_value"`
	IntValue       int               `json:"int_value"`
	Int8Value      int8              `json:"int8_value"`
	Int16Value     int16             `json:"int16_value"`
	Int32Value     int32             `json:"int32_value"`
	Int64Value     int64             `json:"int64_value"`
	UintValue      uint              `json:"uint_value"`
	Uint8Value     uint8             `json:"uint8_value"`
	Uint16Value    uint16            `json:"uint16_value"`
	Uint32Value    uint32            `json:"uint32_value"`
	Uint64Value    uint64            `json:"uint64_value"`
	Float32Value   float32           `json:"float32_value"`
	Float64Value   float64           `json:"float64_value"`
	DateValue      sheetdb.Date      `json:"date_value"`
	DatetimeValue  sheetdb.Datetime  `json:"datetime_value"`
	PBoolValue     *bool             `json:"p_bool_value"`
	PIntValue      *int              `json:"p_int_value"`
	PInt8Value     *int8             `json:"p_int8_value"`
	PInt16Value    *int16            `json:"p_int16_value"`
	PInt32Value    *int32            `json:"p_int32_value"`
	PInt64Value    *int64            `json:"p_int64_value"`
	PUintValue     *uint             `json:"p_uint_value"`
	PUint8Value    *uint8            `json:"p_uint8_value"`
	PUint16Value   *uint16           `json:"p_uint16_value"`
	PUint32Value   *uint32           `json:"p_uint32_value"`
	PUint64Value   *uint64           `json:"p_uint64_value"`
	PFloat32Value  *float32          `json:"p_float32_value"`
	PFloat64Value  *float64          `json:"p_float64_value"`
	PDateValue     *sheetdb.Date     `json:"p_date_value"`
	PDatetimeValue *sheetdb.Datetime `json:"p_datetime_value"`
}
