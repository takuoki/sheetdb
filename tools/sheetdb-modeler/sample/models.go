package sample

import (
	"github.com/takuoki/sheetdb"
)

var dbClient *sheetdb.Client

//go:generate sheetdb-modeler -type=User -children=Foo,Bar -initial=10001
type User struct {
	UserID   int           `json:"user_id",primary`
	Name     string        `json:"name"`
	Email    string        `json:"email",unique`
	Sex      Sex           `json:"sex"`
	Birthday *sheetdb.Date `json:"birthday"`
}

//go:generate sheetdb-modeler -type=Foo -parent=User
type Foo struct {
	UserID int     `json:"user_id",primary`
	FooID  int     `json:"foo_id",primary`
	Value  float32 `json:"value"`
	Note   string  `json:"note",allowempty`
}

//go:generate sheetdb-modeler -type=Bar -parent=User
type Bar struct {
	UserID   int              `json:"user_id",primary`
	Datetime sheetdb.Datetime `json:"datetime",primary`
	Value    float32          `json:"value"`
	Note     string           `json:"note",allowempty`
}
