package sample

import (
	"github.com/takuoki/sheetdb"
)

var dbClient *sheetdb.Client

//go:generate sheetdb-modeler -type=User -children=Foo,Bar -initial=10001
type User struct {
	UserID   int           `json:"user_id"` // primary key
	Name     string        `json:"name"`    // not null
	Email    string        `json:"email"`   // unique not null
	Sex      Sex           `json:"sex"`
	Birthday *sheetdb.Date `json:"birthday"`
}

//go:generate sheetdb-modeler -type=Foo -parent=User
type Foo struct {
	UserID int     `json:"user_id"` // primary key
	FooID  int     `json:"foo_id"`  // primary
	Value  float32 `json:"value"`
	Note   string  `json:"note"` // null
}

//go:generate sheetdb-modeler -type=Bar -parent=User
type Bar struct {
	UserID   int              `json:"user_id"`  // primary key
	Datetime sheetdb.Datetime `json:"datetime"` // primary
	Value    float32          `json:"value"`
	Note     string           `json:"note"` // null
}
