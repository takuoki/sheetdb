package sample

import (
	"github.com/takuoki/sheetdb"
)

var dbClient *sheetdb.Client

//go:generate sheetdb-modeler -type=User -children=Foo,Bar -initial=10001

// User is a struct of user.
type User struct {
	UserID   int           `json:"user_id" db:"primarykey"`
	Name     string        `json:"name"`
	Email    string        `json:"email" db:"unique"`
	Sex      Sex           `json:"sex"`
	Birthday *sheetdb.Date `json:"birthday"`
}

//go:generate sheetdb-modeler -type=Foo -parent=User

// Foo is a struct of foo which is a child of user.
type Foo struct {
	UserID int     `json:"user_id" db:"primarykey"`
	FooID  int     `json:"foo_id" db:"primarykey"`
	Value  float32 `json:"value"`
	Note   string  `json:"note" db:"allowempty"`
}

//go:generate sheetdb-modeler -type=Bar -parent=User

// Bar is a struct of bar which is a child of user.
type Bar struct {
	UserID   int              `json:"user_id" db:"primarykey"`
	Datetime sheetdb.Datetime `json:"datetime" db:"primarykey"`
	Value    float32          `json:"value"`
	Note     string           `json:"note" db:"allowempty"`
}
