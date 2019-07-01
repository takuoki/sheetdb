package main

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type model struct {
	Name                     string
	NamePlural               string
	NameLower                string
	NameLowerPlural          string
	Fields                   []field
	PkNames                  []string
	PkNameLowers             []string
	PkTypes                  []string
	NonPkNameLowers          []string
	NonPkTypes               []string
	FieldNames               []string
	FieldNameLowers          []string
	FieldTypes               []string
	ThisKeyName              string
	ThisKeyNameLower         string
	ThisKeyType              string
	Parent                   *model
	Children                 []model
	ChildrenNames            []string
	ChildrenNamePlurals      []string
	ChildrenNameLowers       []string
	ChildrenNameLowerPlurals []string
}

type option struct {
	ClientName string
	Initial    int
}

type field struct {
	Name              string
	NameLower         string
	Typ               string // *sheetdb.Date
	NonPointerTyp     string // sheetdb.Date
	TypPackage        string // sheetdb
	TypWithoutPackage string // Date
	PointerTyp        bool
	PrimaryKey        bool
	ParentKey         bool
	AllowEmpty        bool
	Unique            bool
}

var (
	sampleUser = model{
		Name:             "User",
		NamePlural:       inflection.Plural("User"),
		NameLower:        strcase.ToLowerCamel("User"),
		NameLowerPlural:  inflection.Plural(strcase.ToLowerCamel("User")),
		PkNames:          []string{"UserID"},
		PkNameLowers:     []string{"userID"},
		PkTypes:          []string{"int"},
		NonPkNameLowers:  []string{"name", "email", "sex", "birthday"},
		NonPkTypes:       []string{"string", "string", "Sex", "*sheetdb.Date"},
		FieldNames:       []string{"UserID", "Name", "Email", "Sex", "Birthday"},
		FieldNameLowers:  []string{"userID", "name", "email", "sex", "birthday"},
		FieldTypes:       []string{"int", "string", "string", "Sex", "*sheetdb.Date"},
		ThisKeyName:      "UserID",
		ThisKeyNameLower: "userID",
		ThisKeyType:      "int",
		Fields: []field{
			{Name: "UserID", NameLower: "userID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true},
			{Name: "Name", NameLower: "name", Typ: "string", NonPointerTyp: "string"},
			{Name: "Email", NameLower: "email", Typ: "string", NonPointerTyp: "string", Unique: true},
			{Name: "Sex", NameLower: "sex", Typ: "Sex", NonPointerTyp: "Sex"},
			{Name: "Birthday", NameLower: "birthday", Typ: "*sheetdb.Date", NonPointerTyp: "sheetdb.Date", TypPackage: "sheetdb", TypWithoutPackage: "Date", PointerTyp: true},
		},
		Children:                 []model{sampleFoo, sampleBar},
		ChildrenNames:            []string{"Foo", "Bar"},
		ChildrenNamePlurals:      []string{"Foos", "Bars"},
		ChildrenNameLowers:       []string{"foo", "bar"},
		ChildrenNameLowerPlurals: []string{"foos", "bars"},
	}

	sampleFoo = model{
		Name:             "Foo",
		NamePlural:       inflection.Plural("Foo"),
		NameLower:        strcase.ToLowerCamel("Foo"),
		NameLowerPlural:  inflection.Plural(strcase.ToLowerCamel("Foo")),
		PkNames:          []string{"UserID", "FooID"},
		PkNameLowers:     []string{"userID", "fooID"},
		PkTypes:          []string{"int", "int"},
		NonPkNameLowers:  []string{"value", "note"},
		NonPkTypes:       []string{"float32", "string"},
		FieldNames:       []string{"UserID", "FooID", "Value", "Note"},
		FieldNameLowers:  []string{"userID", "fooID", "value", "note"},
		FieldTypes:       []string{"int", "int", "float32", "string"},
		ThisKeyName:      "FooID",
		ThisKeyNameLower: "fooID",
		ThisKeyType:      "int",
		Fields: []field{
			{Name: "UserID", NameLower: "userID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true, ParentKey: true},
			{Name: "FooID", NameLower: "fooID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true},
			{Name: "Value", NameLower: "value", Typ: "float32", NonPointerTyp: "float32"},
			{Name: "Note", NameLower: "note", Typ: "string", NonPointerTyp: "string", AllowEmpty: true},
		},
		Parent: &model{
			Name:             "User",
			PkNames:          []string{"UserID"},
			PkNameLowers:     []string{"userID"},
			PkTypes:          []string{"int"},
			NonPkNameLowers:  []string{"name", "email", "sex", "birthday"},
			NonPkTypes:       []string{"string", "string", "Sex", "*sheetdb.Date"},
			FieldNames:       []string{"UserID", "Name", "Email", "Sex", "Birthday"},
			FieldNameLowers:  []string{"userID", "name", "email", "sex", "birthday"},
			FieldTypes:       []string{"int", "string", "string", "Sex", "*sheetdb.Date"},
			ThisKeyName:      "UserID",
			ThisKeyNameLower: "userID",
			ThisKeyType:      "int",
			Fields: []field{
				{Name: "UserID", NameLower: "userID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true},
				{Name: "Name", NameLower: "name", Typ: "string", NonPointerTyp: "string"},
				{Name: "Email", NameLower: "email", Typ: "string", NonPointerTyp: "string", Unique: true},
				{Name: "Sex", NameLower: "sex", Typ: "Sex", NonPointerTyp: "Sex"},
				{Name: "Birthday", NameLower: "birthday", Typ: "*sheetdb.Date", NonPointerTyp: "sheetdb.Date", TypPackage: "sheetdb", TypWithoutPackage: "Date", PointerTyp: true},
			},
		},
	}

	sampleBar = model{
		Name:             "Bar",
		NamePlural:       inflection.Plural("Bar"),
		NameLower:        strcase.ToLowerCamel("Bar"),
		NameLowerPlural:  inflection.Plural(strcase.ToLowerCamel("Bar")),
		PkNames:          []string{"UserID", "Datetime"},
		PkNameLowers:     []string{"userID", "datetime"},
		PkTypes:          []string{"int", "sheetdb.Datetime"},
		NonPkNameLowers:  []string{"value", "note"},
		NonPkTypes:       []string{"float32", "string"},
		FieldNames:       []string{"UserID", "Datetime", "Value", "Note"},
		FieldNameLowers:  []string{"userID", "datetime", "value", "note"},
		FieldTypes:       []string{"int", "sheetdb.Datetime", "float32", "string"},
		ThisKeyName:      "Datetime",
		ThisKeyNameLower: "datetime",
		ThisKeyType:      "sheetdb.Datetime",
		Fields: []field{
			{Name: "UserID", NameLower: "userID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true, ParentKey: true},
			{Name: "Datetime", NameLower: "datetime", Typ: "sheetdb.Datetime", NonPointerTyp: "sheetdb.Datetime", TypPackage: "sheetdb", TypWithoutPackage: "Datetime", PrimaryKey: true},
			{Name: "Value", NameLower: "value", Typ: "float32", NonPointerTyp: "float32"},
			{Name: "Note", NameLower: "note", Typ: "string", NonPointerTyp: "string", AllowEmpty: true},
		},
		Parent: &model{
			Name:             "User",
			PkNames:          []string{"UserID"},
			PkNameLowers:     []string{"userID"},
			PkTypes:          []string{"int"},
			NonPkNameLowers:  []string{"name", "email", "sex", "birthday"},
			NonPkTypes:       []string{"string", "string", "Sex", "*sheetdb.Date"},
			FieldNames:       []string{"UserID", "Name", "Email", "Sex", "Birthday"},
			FieldNameLowers:  []string{"userID", "name", "email", "sex", "birthday"},
			FieldTypes:       []string{"int", "string", "string", "Sex", "*sheetdb.Date"},
			ThisKeyName:      "UserID",
			ThisKeyNameLower: "userID",
			ThisKeyType:      "int",
			Fields: []field{
				{Name: "UserID", NameLower: "userID", Typ: "int", NonPointerTyp: "int", PrimaryKey: true},
				{Name: "Name", NameLower: "name", Typ: "string", NonPointerTyp: "string"},
				{Name: "Email", NameLower: "email", Typ: "string", NonPointerTyp: "string", Unique: true},
				{Name: "Sex", NameLower: "sex", Typ: "Sex", NonPointerTyp: "Sex"},
				{Name: "Birthday", NameLower: "birthday", Typ: "*sheetdb.Date", NonPointerTyp: "sheetdb.Date", TypPackage: "sheetdb", TypWithoutPackage: "Date", PointerTyp: true},
			},
		},
	}
)

func (g *Generator) generate(typeName string) {
	m := sampleUser
	switch typeName {
	case "Foo":
		m = sampleFoo
	case "Bar":
		m = sampleBar
	}
	g.output(m, option{
		ClientName: "dbClient",
		Initial:    10001,
	})
}
