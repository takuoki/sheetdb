package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/takuoki/clmconv"
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

func xxxfixes(strs []string, prefix, suffix string) []string {
	r := []string{}
	for _, str := range strs {
		r = append(r, prefix+str+suffix)
	}
	return r
}

func join(s1, s2 []string, delimiter, separater string) string {
	if len(s1) != len(s2) {
		panic("join functions arguments are invalid: the length of s1 and s2 should be same")
	}
	r := []string{}
	for i := 0; i < len(s1); i++ {
		r = append(r, s1[i]+delimiter+s2[i])
	}
	return strings.Join(r, separater)
}

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

func (g *Generator) output(m model, o option) {
	g.outputImport(m)
	g.outputConst(m)
	g.outputVar(m)
	g.outputInit(m, o)
	g.outputLoad(m)
	g.outputGet(m)
	g.outputGetList(m)
	g.outputAdd(m, o)
	g.outputUpdate(m)
	g.outputDelete(m)
	g.outputValidate(m)
	g.outputParse(m)
	g.outputAsync(m, o)
}

func (g *Generator) outputImport(m model) {

	var includeNumber bool
	for _, f := range m.Fields {
		switch f.Typ {
		case "int", "float32": // TODO
			includeNumber = true
			break
		}
	}

	g.Printf("import (\n")
	if includeNumber {
		g.Printf("\t\"strconv\"\n")
	}
	g.Printf("\t\"sync\"\n")
	g.Printf("\t\"time\"\n")
	g.Printf("\n")
	g.Printf("\t\"github.com/takuoki/gsheets\"\n")
	g.Printf("\t\"github.com/takuoki/sheetdb\"\n")
	g.Printf(")\n\n")
}

func (g *Generator) outputConst(m model) {
	g.Printf("const (\n")
	g.Printf("\t_%s_sheetName = \"%s\"\n", m.Name, m.NameLowerPlural)
	for i, f := range m.Fields {
		g.Printf("\t_%s_column_%s = %d // %s\n", m.Name, f.Name, i, clmconv.Itoa(i))
	}
	g.Printf("\t_%s_column_UpdatedAt = %d // %s\n", m.Name, len(m.Fields), clmconv.Itoa(len(m.Fields)))
	g.Printf("\t_%s_column_DeletedAt = %d // %s\n", m.Name, len(m.Fields)+1, clmconv.Itoa(len(m.Fields)+1))
	g.Printf(")\n\n")
}

func (g *Generator) outputVar(m model) {
	// TODO: comment
	g.Printf("var (\n")
	g.Printf("\t_%s_mutex = sync.RWMutex{}\n", m.Name)
	g.Printf("\t_%[1]s_cache = map[%[2]s]*%[1]s{}\n", m.Name, strings.Join(m.PkTypes, "]map["))
	g.Printf("\t_%s_rowNoMap = map[%s]int{}\n", m.Name, strings.Join(m.PkTypes, "]map["))
	g.Printf("\t_%s_maxRowNo int\n", m.Name)
	g.Printf(")\n\n")
}

func (g *Generator) outputInit(m model, o option) {
	g.Printf("func init() {\n")
	g.Printf("\t%[1]s.AddModel(\"%[2]s\", _%[2]s_load)\n", o.ClientName, m.Name)
	g.Printf("}\n\n")
}

func (g *Generator) outputLoad(m model) {
	g.Printf("func _%s_load(data *gsheets.Sheet) error {\n\n", m.Name)

	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n\n", m.Name)

	g.Printf("\t_%[1]s_cache = map[%[2]s]*%[1]s{}\n", m.Name, strings.Join(m.PkTypes, "]map["))
	g.Printf("\t_%s_rowNoMap = map[%s]int{}\n", m.Name, strings.Join(m.PkTypes, "]map["))
	g.Printf("\t_%s_maxRowNo = 0\n\n", m.Name)

	g.Printf("\tfor i, r := range data.Rows() {\n")
	g.Printf("\t\tif i == 0 {\n")
	g.Printf("\t\t\tcontinue\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\tif r.Value(_%s_column_DeletedAt) != \"\" {\n", m.Name)
	g.Printf("\t\t\t_%s_maxRowNo++\n", m.Name)
	g.Printf("\t\t\tcontinue\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\tif r.Value(_%s_column_%s) == \"\" {\n", m.Name, m.Fields[0].Name)
	g.Printf("\t\t\tbreak\n")
	g.Printf("\t\t}\n\n")

	for _, f := range m.Fields {
		if f.Typ == "string" {
			g.Printf("\t\t%[3]s := r.Value(_%[1]s_column_%[2]s)\n", m.Name, f.Name, f.NameLower)
			if f.Unique {
				g.Printf("\t\tif err := _%[1]s_validate%[2]s(%[3]s, nil); err != nil {\n", m.Name, f.Name, f.NameLower)
			} else {
				g.Printf("\t\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLower)
			}
		} else {
			g.Printf("\t\t%[3]s, err := _%[1]s_parse%[2]s(r.Value(_%[1]s_column_%[2]s))\n", m.Name, f.Name, f.NameLower)
			g.Printf("\t\tif err != nil {\n")
		}
		g.Printf("\t\t\treturn err\n")
		g.Printf("\t\t}\n")
	}
	g.Printf("\n")

	g.Printf("\t\t%[2]s := %[1]s{\n", m.Name, m.NameLower)
	for _, f := range m.Fields {
		g.Printf("\t\t\t%s: %s,\n", f.Name, f.NameLower)
	}
	g.Printf("\t\t}\n\n")

	g.Printf("\t\t_%s_maxRowNo++\n", m.Name)
	g.outputParentMap(m)
	g.Printf("\t\t_%[1]s_cache[%[3]s] = &%[2]s\n", m.Name, m.NameLower, strings.Join(m.PkNameLowers, "]["))
	g.Printf("\t\t_User_rowNoMap[userID] = _User_maxRowNo\n")
	g.Printf("\t}\n\n")

	g.Printf("\treturn nil\n")
	g.Printf("}\n\n")
}

func (g *Generator) outputParentMap(m model) {
	for i, f := range m.PkNames {
		if f == m.ThisKeyName {
			break
		}
		g.Printf("if _, ok := _%s_cache", m.Name)
		for j, f2 := range m.PkNames {
			if j > i {
				break
			}
			g.Printf("[%s.%s]", m.NameLower, f2)
		}
		g.Printf("; !ok {\n")
		g.Printf("\t	_%[1]s_cache", m.Name)
		for j, f2 := range m.PkNames {
			if j > i {
				break
			}
			g.Printf("[%s.%s]", m.NameLower, f2)
		}
		g.Printf(" = map")
		for j, f2 := range m.PkTypes {
			if j <= i {
				continue
			}
			g.Printf("[%s]", f2)
		}
		g.Printf("*%[1]s{}\n", m.Name)
		g.Printf("}\n")
	}
}

func (g *Generator) outputGet(m model) {
	g.Printf("func Get%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, join(m.PkNameLowers, m.PkTypes, " ", ", "))
	g.Printf("\t_%s_mutex.RLock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.RUnlock()\n", m.Name)
	g.Printf("\tif v, ok := _%s_cache[%s]; ok {\n", m.Name, strings.Join(m.PkNameLowers, "]["))
	g.Printf("\t\treturn v, nil\n")
	g.Printf("\t}\n")
	g.Printf("\treturn nil, &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("}\n\n")
}

func (g *Generator) outputGetList(m model) {
	g.Printf("type %sQuery struct {\n", m.Name)
	g.Printf("\tfilter func(%[2]s *%[1]s) bool\n", m.Name, m.NameLower)
	g.Printf("\tsort   func(%[2]s []*%[1]s)\n", m.Name, m.NameLowerPlural)
	g.Printf("}\n\n")

	g.Printf("type %[1]sQueryOption func(query *%[1]sQuery) *%[1]sQuery\n\n", m.Name)

	g.Printf("func %[1]sFilter(filterFunc func(%[2]s *%[1]s) bool) func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name, m.NameLower)
	g.Printf("\treturn func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name)
	g.Printf("\t\tif query != nil {\n")
	g.Printf("\t\t\tquery.filter = filterFunc\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\treturn query\n")
	g.Printf("\t}\n")
	g.Printf("}\n\n")

	g.Printf("func %[1]sSort(sortFunc func(%[2]s []*%[1]s)) func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name, m.NameLowerPlural)
	g.Printf("\treturn func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name)
	g.Printf("\t\tif query != nil {\n")
	g.Printf("\t\t\tquery.sort = sortFunc\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\treturn query\n")
	g.Printf("\t}\n")
	g.Printf("}\n\n")

	if m.Parent == nil {
		g.Printf("func Get%[2]s(opts ...%[1]sQueryOption) ([]*%[1]s, error) {\n", m.Name, m.NamePlural)
	} else {
		g.Printf("func Get%[2]s(%[3]s, opts ...%[1]sQueryOption) ([]*%[1]s, error) {\n", m.Name, m.NamePlural, join(m.Parent.PkNameLowers, m.Parent.PkTypes, " ", ", "))
	}

	g.Printf("\t%[2]sQuery := &%[1]sQuery{}\n", m.Name, m.NameLower)
	g.Printf("\tfor _, opt := range opts {\n")
	g.Printf("\t\t%[1]sQuery = opt(%[1]sQuery)\n", m.NameLower)
	g.Printf("\t}\n")
	g.Printf("\t_%s_mutex.RLock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.RUnlock()\n", m.Name)
	g.Printf("\t%[2]s := []*%[1]s{}\n", m.Name, m.NameLowerPlural)
	g.Printf("\tif %sQuery.filter != nil {\n", m.NameLower)

	if m.Parent == nil {
		g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
	} else {
		g.Printf("\t\tfor _, v := range _%s_cache[%s] {\n", m.Name, strings.Join(m.Parent.PkNameLowers, "]["))
	}

	g.Printf("\t\t\tif %sQuery.filter(v) {\n", m.NameLower)
	g.Printf("\t\t\t\t%[1]s = append(%[1]s, v)\n", m.NameLowerPlural)
	g.Printf("\t\t\t}\n")
	g.Printf("\t\t}\n")
	g.Printf("\t} else {\n")

	if m.Parent == nil {
		g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
	} else {
		g.Printf("\t\tfor _, v := range _%s_cache[%s] {\n", m.Name, strings.Join(m.Parent.PkNameLowers, "]["))
	}

	g.Printf("\t\t\t%[1]s = append(%[1]s, v)\n", m.NameLowerPlural)
	g.Printf("\t\t}\n")
	g.Printf("\t}\n")
	g.Printf("\tif %sQuery.sort != nil {\n", m.NameLower)
	g.Printf("\t\t%sQuery.sort(%s)\n", m.NameLower, m.NameLowerPlural)
	g.Printf("\t}\n")
	g.Printf("\treturn %s, nil\n", m.NameLowerPlural)
	g.Printf("}\n\n")
}

func (g *Generator) outputAdd(m model, o option) {

	if m.Parent == nil {
		if m.ThisKeyType == "int" && m.ThisKeyName[len(m.ThisKeyName)-2:] == "ID" {
			g.Printf("func Add%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
		} else {
			g.Printf("func Add%[1]s(%[2]s %[3]s, %[4]s) (*%[1]s, error) {\n", m.Name, m.ThisKeyNameLower, m.ThisKeyType, join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
		}
	} else {
		if m.ThisKeyType == "int" && m.ThisKeyName[len(m.ThisKeyName)-2:] == "ID" {
			g.Printf("func (m *%[2]s) Add%[1]s(%[3]s) (*%[1]s, error) {\n", m.Name, m.Parent.Name, join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
		} else {
			g.Printf("func (m *%[4]s) Add%[1]s(%[2]s %[3]s, %[5]s) (*%[1]s, error) {\n", m.Name, m.ThisKeyNameLower, m.ThisKeyType, m.Parent.Name, join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
		}
	}

	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, f := range m.Fields {
		if !f.PrimaryKey && f.Typ == "string" {
			if f.Unique {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s, nil); err != nil {\n", m.Name, f.Name, f.NameLower)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")
			} else {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLower)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")
			}
		}
	}

	g.Printf("\t%[2]s := &%[1]s{\n", m.Name, m.NameLower)

	for _, f := range m.Fields {
		if f.PrimaryKey {
			if f.ParentKey {
				g.Printf("\t\t%[1]s: m.%[1]s,\n", f.Name)
			} else if f.Typ != "int" || f.Name[len(f.Name)-2:] != "ID" {
				g.Printf("\t\t%s: %s,\n", f.Name, f.NameLower)
			} else {
				g.Printf("\t\t%[2]s: _%[1]s_maxRowNo + %[3]d,\n", m.Name, f.Name, o.Initial)
			}
		} else {
			g.Printf("\t\t%s: %s,\n", f.Name, f.NameLower)
		}
	}

	g.Printf("\t}\n")
	g.Printf("\tif err := %s._asyncUpdate(); err != nil {\n", m.NameLower)
	g.Printf("\t\treturn nil, err\n")
	g.Printf("\t}\n")
	g.Printf("\t_%s_maxRowNo++\n", m.Name)

	g.outputParentMap(m)

	g.Printf("\t_%[1]s_cache[%[3]s] = %[2]s\n", m.Name, m.NameLower, strings.Join(xxxfixes(m.PkNames, m.NameLower+".", ""), "]["))
	g.Printf("\t_%[1]s_rowNoMap[%[2]s] = _%[1]s_maxRowNo\n", m.Name, strings.Join(xxxfixes(m.PkNames, m.NameLower+".", ""), "]["))
	g.Printf("\treturn %s, nil\n", m.NameLower)
	g.Printf("}\n\n")

	if m.Parent != nil {
		if m.ThisKeyType == "int" && m.ThisKeyName[len(m.ThisKeyName)-2:] == "ID" {
			g.Printf("func Add%[1]s(%[2]s, %[3]s) (*%[1]s, error) {\n", m.Name, join(m.Parent.PkNameLowers, m.Parent.PkTypes, " ", ", "), join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
		} else {
			g.Printf("func Add%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, join(m.FieldNameLowers, m.FieldTypes, " ", ", "))
		}
		g.outputGetParent(*m.Parent, true, 0)
		if m.ThisKeyType == "int" && m.ThisKeyName[len(m.ThisKeyName)-2:] == "ID" {
			g.Printf("\treturn m.Add%s(%s)\n", m.Name, strings.Join(m.NonPkNameLowers, ", "))
		} else {
			g.Printf("\treturn m.Add%s(%s, %s)\n", m.Name, m.ThisKeyNameLower, strings.Join(m.NonPkNameLowers, ", "))
		}
		g.Printf("}\n\n")
	}
}

func (g *Generator) outputGetParent(m model, returnNil bool, i int) {
	if m.Parent != nil {
		g.outputGetParent(*m.Parent, returnNil, i+1)
	}
	if i == 0 {
		g.Printf("\tm, err := Get%[1]s(%[2]s)\n", m.Name, strings.Join(m.PkNameLowers, ", "))
	} else {
		g.Printf("\tm%[3]d, err := Get%[1]s(%[2]s)\n", m.Name, strings.Join(m.PkNameLowers, ", "), i)
	}
	g.Printf("\tif err != nil {\n")
	if returnNil {
		g.Printf("\t\treturn nil, err\n")
	} else {
		g.Printf("\t\treturn err\n")
	}
	g.Printf("\t}\n")
}

func (g *Generator) outputUpdate(m model) {

	if m.Parent == nil {
		g.Printf("func Update%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, join(m.FieldNameLowers, m.FieldTypes, " ", ", "))
	} else {
		g.Printf("func (m *%[2]s) Update%[1]s(%[3]s %[4]s, %[5]s) (*%[1]s, error) {\n", m.Name, m.Parent.Name, m.ThisKeyNameLower, m.ThisKeyType, join(m.NonPkNameLowers, m.NonPkTypes, " ", ", "))
	}

	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, f := range m.Fields {
		if !f.PrimaryKey && f.Typ == "string" {
			if f.Unique {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s, &%[4]s); err != nil {\n", m.Name, f.Name, f.NameLower, strings.Join(m.PkNameLowers, ", &"))
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")
			} else {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLower)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")
			}
		}
	}

	if m.Parent == nil {
		g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s]\n", m.Name, m.NameLower, strings.Join(m.PkNameLowers, "]["))
	} else {
		g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s][%[4]s]\n", m.Name, m.NameLower, strings.Join(xxxfixes(m.Parent.PkNames, "m.", ""), "]["), m.ThisKeyNameLower)
	}

	g.Printf("\tif !ok {\n")
	g.Printf("\t\treturn nil, &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("\t}\n")
	g.Printf("\t%[1]sCopy := *%[1]s\n", m.NameLower)

	for _, f := range m.Fields {
		if !f.PrimaryKey {
			g.Printf("\t%sCopy.%s = %s\n", m.NameLower, f.Name, f.NameLower)
		}
	}

	g.Printf("\tif err := (&%sCopy)._asyncUpdate(); err != nil {\n", m.NameLower)
	g.Printf("\t\treturn nil, err\n")
	g.Printf("\t}\n")
	g.Printf("\t%[1]s = &%[1]sCopy\n", m.NameLower)
	g.Printf("\treturn %s, nil\n", m.NameLower)
	g.Printf("}\n\n")

	if m.Parent != nil {
		g.Printf("func Update%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, join(m.FieldNameLowers, m.FieldTypes, " ", ", "))
		g.outputGetParent(*m.Parent, true, 0)
		g.Printf("\treturn m.Update%s(%s, %s)\n", m.Name, m.ThisKeyNameLower, strings.Join(m.NonPkNameLowers, ", "))
		g.Printf("}\n\n")
	}
}

func (g *Generator) outputDelete(m model) {

	if m.Parent == nil {
		g.Printf("func Delete%[1]s(%[2]s %[3]s) error {\n", m.Name, m.ThisKeyNameLower, m.ThisKeyType)
	} else {
		g.Printf("func (m *%[2]s) Delete%[1]s(%[3]s %[4]s) error {\n", m.Name, m.Parent.Name, m.ThisKeyNameLower, m.ThisKeyType)
	}

	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, child := range m.Children {
		g.Printf("\t_%s_mutex.Lock()\n", child.Name)
		g.Printf("\tdefer _%s_mutex.Unlock()\n", child.Name)
	}

	if m.Parent == nil {
		g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s]\n", m.Name, m.NameLower, m.ThisKeyNameLower)
	} else {
		g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s][%[4]s]\n", m.Name, m.NameLower, strings.Join(xxxfixes(m.Parent.PkNames, "m.", ""), "]["), m.ThisKeyNameLower)
	}

	g.Printf("\tif !ok {\n")
	g.Printf("\t\treturn &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("\t}\n")

	for _, child := range m.Children {
		g.Printf("\tvar %[2]s []*%[1]s\n", child.Name, child.NameLowerPlural)
		g.Printf("\tfor _, v := range _%s_cache[%s] {\n", child.Name, strings.Join(m.PkNameLowers, "]["))
		g.Printf("\t\t%[1]s = append(%[1]s, v)\n", child.NameLowerPlural)
		g.Printf("\t}\n")
	}

	g.Printf("\tif err := %s._asyncDelete(%s); err != nil {\n", m.NameLower, strings.Join(m.ChildrenNameLowerPlurals, ", "))
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")

	if m.Parent == nil {
		g.Printf("\tdelete(_%[1]s_cache, %[2]s)\n", m.Name, m.ThisKeyNameLower)
	} else {
		g.Printf("\tdelete(_%[1]s_cache[%[3]s], %[2]s)\n", m.Name, m.ThisKeyNameLower, strings.Join(xxxfixes(m.Parent.PkNames, "m.", ""), "]["))
	}

	for _, child := range m.Children {
		if m.Parent == nil {
			g.Printf("\tdelete(_%[1]s_cache, %[2]s)\n", child.Name, m.ThisKeyNameLower)
		} else {
			g.Printf("\tdelete(_%[1]s_cache[%[3]s], %[2]s)\n", child.Name, m.ThisKeyNameLower, strings.Join(xxxfixes(m.Parent.PkNames, "m.", ""), "]["))
		}
	}

	g.Printf("\treturn nil\n")
	g.Printf("}\n\n")

	if m.Parent != nil {
		g.Printf("func Delete%[1]s(%[2]s) error {\n", m.Name, join(m.PkNameLowers, m.PkTypes, " ", ", "))
		g.outputGetParent(*m.Parent, false, 0)
		g.Printf("\treturn m.Delete%s(%s)\n", m.Name, m.ThisKeyNameLower)
		g.Printf("}\n\n")
	}
}

func (g *Generator) outputValidate(m model) {
	// TODO: use cache for unique check
	pkNameAndTypes := []string{}
	pkEqualConditions := []string{}
	for _, f := range m.Fields {
		if f.PrimaryKey {
			pkNameAndTypes = append(pkNameAndTypes, f.NameLower+" *"+f.Typ)
			pkEqualConditions = append(pkEqualConditions, "v."+f.Name+" == *"+f.NameLower)
		}
	}
	for _, f := range m.Fields {
		if f.Typ != "string" {
			continue
		}
		if f.Unique {
			g.Printf("func _%[1]s_validate%[2]s(%[3]s string, %[4]s) error {\n", m.Name, f.Name, f.NameLower, strings.Join(pkNameAndTypes, ", "))
		} else {
			g.Printf("func _%[1]s_validate%[2]s(%[3]s string) error {\n", m.Name, f.Name, f.NameLower)
		}

		if !f.AllowEmpty {
			g.Printf("\tif %s == \"\" {\n", f.NameLower)
			g.Printf("\t\treturn &sheetdb.EmptyStringError{FieldName: \"%s\"}\n", f.Name)
			g.Printf("\t}\n")
		}

		if f.Unique {
			g.Printf("\tif %s == nil {\n", strings.Join(m.PkNameLowers, " == nil || "))
			g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
			g.Printf("\t\t\tif %[2]s == v.%[1]s {\n", f.Name, f.NameLower)
			g.Printf("\t\t\t\treturn &sheetdb.DuplicationError{FieldName: \"%s\"}\n", f.Name)
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t}\n")
			g.Printf("\t} else {\n")
			g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
			g.Printf("\t\t\tif %s {\n", strings.Join(pkEqualConditions, " && "))
			g.Printf("\t\t\t\tcontinue\n")
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t\tif %[2]s == v.%[1]s {\n", f.Name, f.NameLower)
			g.Printf("\t\t\t\treturn &sheetdb.DuplicationError{FieldName: \"%s\"}\n", f.Name)
			g.Printf("\t\t\t}\n")
			g.Printf("\t\t}\n")
			g.Printf("\t}\n")
		}

		g.Printf("\treturn nil\n")
		g.Printf("}\n\n")
	}
}

func (g *Generator) outputParse(m model) {

	for _, f := range m.Fields {
		if f.Typ == "string" {
			continue
		}
		g.Printf("func _%s_parse%s(%s string) (%s, error) {\n", m.Name, f.Name, f.NameLower, f.Typ)

		if f.PointerTyp {
			g.Printf("\tvar val %s\n", f.Typ)
			g.Printf("\tif %s != \"\" {\n", f.NameLower)
		}

		switch f.Typ {
		case "int":
			g.Printf("\tv, err := strconv.Atoi(%s)\n", f.NameLower)
		case "float32":
			g.Printf("\tv, err := strconv.ParseFloat(%s, 32)\n", f.NameLower)
		default:
			if f.TypPackage == "" {
				g.Printf("\tv, err := New%s(%s)\n", f.NonPointerTyp, f.NameLower)
			} else {
				g.Printf("\tv, err := %s.New%s(%s)\n", f.TypPackage, f.TypWithoutPackage, f.NameLower)
			}
		}

		g.Printf("\tif err != nil {\n")
		g.Printf("\t\treturn %[2]s, &sheetdb.InvalidValueError{FieldName: \"%[1]s\", Err: err}\n", f.Name, nilValue(f.Typ))
		g.Printf("\t}\n")

		if f.PointerTyp {
			g.Printf("\t\tval = &v\n")
			g.Printf("\t}\n")
			g.Printf("\treturn val, nil\n")
		} else {
			switch f.Typ {
			case "float32":
				g.Printf("\treturn float32(v), nil\n")
			default:
				g.Printf("\treturn v, nil\n")
			}
		}

		g.Printf("}\n\n")
	}
}

func (g *Generator) outputAsync(m model, o option) {

	g.Printf("func (m *%s) _asyncUpdate() error {\n", m.Name)
	g.Printf("\tdata := []gsheets.UpdateValue{\n")
	g.Printf("\t\t{\n")
	g.Printf("\t\t\tSheetName: _%s_sheetName,\n", m.Name)
	g.Printf("\t\t\tRowNo:     _%[1]s_rowNoMap[%[2]s],\n", m.Name, strings.Join(xxxfixes(m.PkNames, "m.", ""), "]["))
	g.Printf("\t\t\tValues: []interface{}{\n")

	for _, f := range m.Fields {
		switch f.Typ {
		case "int", "float32", "string":
			g.Printf("\t\t\t\tm.%s,\n", f.Name)
		default:
			g.Printf("\t\t\t\tm.%s.String(),\n", f.Name)
		}
	}

	g.Printf("\t\t\t\ttime.Now(),\n")
	g.Printf("\t\t\t\t\"\",\n")
	g.Printf("\t\t\t},\n")
	g.Printf("\t\t},\n")
	g.Printf("\t}\n")
	g.Printf("\treturn %s.AsyncUpdate(data)\n", o.ClientName)
	g.Printf("}\n\n")

	g.Printf("func (m *%[1]s) _asyncDelete(%[2]s) error {\n", m.Name, join(m.ChildrenNameLowerPlurals, m.ChildrenNames, " []*", ", "))
	g.Printf("\tnow := time.Now()\n")
	g.Printf("\tdata := []gsheets.UpdateValue{\n")
	g.Printf("\t\t{\n")
	g.Printf("\t\t\tSheetName: _%s_sheetName,\n", m.Name)
	g.Printf("\t\t\tRowNo:     _%[1]s_rowNoMap[%[2]s],\n", m.Name, strings.Join(xxxfixes(m.PkNames, "m.", ""), "]["))
	g.Printf("\t\t\tValues: []interface{}{\n")

	for _, f := range m.Fields {
		switch f.Typ {
		case "int", "float32", "string":
			g.Printf("\t\t\t\tm.%s,\n", f.Name)
		default:
			g.Printf("\t\t\t\tm.%s.String(),\n", f.Name)
		}
	}

	g.Printf("\t\t\t\tnow,\n")
	g.Printf("\t\t\t\tnow,\n")
	g.Printf("\t\t\t},\n")
	g.Printf("\t\t},\n")
	g.Printf("\t}\n")

	for i, child := range m.Children {
		g.Printf("\tfor _, v := range %s {\n", child.NameLowerPlural)
		g.Printf("\t\tdata = append(data, gsheets.UpdateValue{\n")
		g.Printf("\t\t\tSheetName: _%s_sheetName,\n", child.Name)
		g.Printf("\t\t\tRowNo:     _%s_rowNoMap[%s],\n", child.Name, strings.Join(xxxfixes(m.Children[i].PkNames, "v.", ""), "]["))
		g.Printf("\t\t\tValues: []interface{}{\n")

		for _, f := range child.Fields {
			switch f.Typ {
			case "int", "float32", "string":
				g.Printf("\t\t\t\tv.%s,\n", f.Name)
			default:
				g.Printf("\t\t\t\tv.%s.String(),\n", f.Name)
			}
		}

		g.Printf("\t\t\t\tnow,\n")
		g.Printf("\t\t\t\tnow,\n")
		g.Printf("\t\t\t},\n")
		g.Printf("\t\t})\n")
		g.Printf("\t}\n")
	}

	g.Printf("\treturn %s.AsyncUpdate(data)\n", o.ClientName)
	g.Printf("}\n")
}

func nilValue(t string) string {
	if t == "" {
		return ""
	}
	if t[0] == '*' {
		return "nil"
	}
	switch t {
	case "int", "float32", "float64":
		return "0"
	}
	// TODO
	return "v"
}
