package main

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/takuoki/clmconv"
)

type model struct {
	Name     string
	Fields   []field
	Parent   *model
	Children []model
	NamePlural string
	NameLowerCamel string
	NameLowerCamelPlural string
}

type option struct {
	ClientName string
	Initial int
}

func (m model) PrimaryKeys() []field {
	var ks []field
	for _, f := range m.Fields {
		if f.PrimaryKey {
			ks = append(ks, f)
		}
	}
	return ks
}

type field struct {
	Name       string
	NameLowerCamel       string
	Typ        string
	PrimaryKey bool
	AllowEmpty bool
	Unique     bool
	// TODO: if child, which are parent key
}

func sampleModel() model {
	return model{
		Name: "User",
		NamePlural: inflection.Plural("User"),
		NameLowerCamel: strcase.ToLowerCamel("User"),
		NameLowerCamelPlural: inflection.Plural(strcase.ToLowerCamel("User")),
		Fields: []field{
			{Name: "UserID", NameLowerCamel: "userID", Typ: "int", PrimaryKey: true},
			{Name: "Name", NameLowerCamel: "name", Typ: "string"},
			{Name: "Email", NameLowerCamel: "email", Typ: "string", Unique: true},
			{Name: "Sex", NameLowerCamel: "sex", Typ: "Sex"},
			{Name: "Birthday", NameLowerCamel: "birthday", Typ: "*sheetdb.Date"},
		},
		Children: []model{
			model{
				Name: "Foo",
				NamePlural: inflection.Plural("Foo"),
				NameLowerCamel: strcase.ToLowerCamel("Foo"),
				NameLowerCamelPlural: inflection.Plural(strcase.ToLowerCamel("Foo")),		
				Fields: []field{
					{Name: "UserID", NameLowerCamel: "userID", Typ: "int", PrimaryKey: true},
					{Name: "FooID", NameLowerCamel: "fooID", Typ: "int", PrimaryKey: true},
					{Name: "Value", NameLowerCamel: "value", Typ: "float32"},
					{Name: "Note", NameLowerCamel: "note", Typ: "string", AllowEmpty: true},
				},
			},
			model{
				Name: "Bar",
				NamePlural: inflection.Plural("Bar"),
				NameLowerCamel: strcase.ToLowerCamel("Bar"),
				NameLowerCamelPlural: inflection.Plural(strcase.ToLowerCamel("Bar")),		
				Fields: []field{
					{Name: "UserID", NameLowerCamel: "userID", Typ: "int", PrimaryKey: true},
					{Name: "Datetime", NameLowerCamel: "datetime", Typ: "sheetdb.Datetime", PrimaryKey: true},
					{Name: "Value",NameLowerCamel: "value",  Typ: "float32"},
					{Name: "Note", NameLowerCamel: "note", Typ: "string", AllowEmpty: true},
				},
			},
		},
	}
}

func (g *Generator) generate(typeName string) {
	g.output(sampleModel(), option{
		ClientName: "dbClient",
		Initial: 10001,
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
	g.Printf("\t_%s_sheetName = \"%s\"\n", m.Name, m.NameLowerCamelPlural)
	for i, f := range m.Fields {
		g.Printf("\t_%s_column_%s = %d // %s\n", m.Name, f.Name, i, clmconv.Itoa(i))
	}
	g.Printf("\t_%s_column_UpdatedAt = %d // %s\n", m.Name, len(m.Fields), clmconv.Itoa(len(m.Fields)))
	g.Printf("\t_%s_column_DeletedAt = %d // %s\n", m.Name, len(m.Fields)+1, clmconv.Itoa(len(m.Fields)+1))
	g.Printf(")\n\n")
}

func (g *Generator) outputVar(m model) {
	// TODO: comment
	pkTypes := []string{}
	for _, f := range m.PrimaryKeys() {
		pkTypes = append(pkTypes, f.Typ)
	}
	g.Printf("var (\n")
	g.Printf("\t_%s_mutex = sync.RWMutex{}\n", m.Name)
	g.Printf("\t_%[1]s_cache = map[%[2]s]*%[1]s{}\n", m.Name, strings.Join(pkTypes, "]["))
	g.Printf("\t_%s_rowNoMap = map[%s]int{}\n", m.Name, strings.Join(pkTypes, "]["))
	g.Printf("\t_%s_maxRowNo int\n", m.Name)
	g.Printf(")\n\n")
}

func (g *Generator) outputInit(m model, o option) {
	g.Printf("func init() {\n")
	g.Printf("\t%[1]s.AddModel(\"%[2]s\", _%[2]s_load)\n", o.ClientName, m.Name)
	g.Printf("}\n\n")
}

func (g *Generator) outputLoad(m model) {
	pkNames := []string{}
	pkTypes := []string{}
	for _, f := range m.PrimaryKeys() {
		pkNames = append(pkNames, f.NameLowerCamel)
		pkTypes = append(pkTypes, f.Typ)
	}
	g.Printf("func _%s_load(data *gsheets.Sheet) error {\n\n", m.Name)

	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n\n", m.Name)

	g.Printf("\t_%[1]s_cache = map[%[2]s]*%[1]s{}\n", m.Name, strings.Join(pkTypes, "]["))
	g.Printf("\t_%s_rowNoMap = map[%s]int{}\n", m.Name, strings.Join(pkTypes, "]["))
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
			g.Printf("\t\t%[3]s := r.Value(_%[1]s_column_%[2]s)\n", m.Name, f.Name, f.NameLowerCamel)
			if f.Unique {
				g.Printf("\t\tif err := _%[1]s_validate%[2]s(%[3]s, nil); err != nil {\n", m.Name, f.Name, f.NameLowerCamel)
			} else {
				g.Printf("\t\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLowerCamel)
			}
		} else {
			g.Printf("\t\t%[3]s, err := _%[1]s_parse%[2]s(r.Value(_%[1]s_column_%[2]s))\n", m.Name, f.Name, f.NameLowerCamel)
			g.Printf("\t\tif err != nil {\n")
		}
		g.Printf("\t\t\treturn err\n")
		g.Printf("\t\t}\n")
	}
	g.Printf("\n")

	g.Printf("\t\t_%s_maxRowNo++\n", m.Name)
	g.Printf("\t\t_%[1]s_cache[%[2]s] = &%[1]s{\n", m.Name, strings.Join(pkNames, "]["))
	for _, f := range m.Fields {
		g.Printf("\t\t\t%s: %s,\n", f.Name, f.NameLowerCamel)
	}

	g.Printf("\t\t}\n")
	g.Printf("\t\t_User_rowNoMap[userID] = _User_maxRowNo\n")
	g.Printf("\t}\n\n")

	g.Printf("\treturn nil\n")
	g.Printf("}\n\n")
}

func (g *Generator) outputGet(m model) {
	pkNames := []string{}
	pkNameAndTypes := []string{}
	for _, f := range m.PrimaryKeys() {
		pkNames = append(pkNames, f.NameLowerCamel)
		pkNameAndTypes = append(pkNameAndTypes, f.NameLowerCamel+" "+f.Typ)
	}
	g.Printf("func Get%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, strings.Join(pkNameAndTypes, ", "))
	g.Printf("\t_%s_mutex.RLock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.RUnlock()\n", m.Name)
	g.Printf("\tif v, ok := _%s_cache[%s]; ok {\n", m.Name, strings.Join(pkNames, "]["))
	g.Printf("\t\treturn v, nil\n")
	g.Printf("\t}\n")
	g.Printf("\treturn nil, &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("}\n\n")
}

func (g *Generator) outputGetList(m model) {
	g.Printf("type %sQuery struct {\n", m.Name)
	g.Printf("\tfilter func(%[2]s *%[1]s) bool\n", m.Name, m.NameLowerCamel)
	g.Printf("\tsort   func(%[2]s []*%[1]s)\n", m.Name, m.NameLowerCamelPlural)
	g.Printf("}\n\n")

	g.Printf("type %[1]sQueryOption func(query *%[1]sQuery) *%[1]sQuery\n\n", m.Name)

	g.Printf("func %[1]sFilter(filterFunc func(%[2]s *%[1]s) bool) func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name, m.NameLowerCamel)
	g.Printf("\treturn func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name)
	g.Printf("\t\tif query != nil {\n")
	g.Printf("\t\t\tquery.filter = filterFunc\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\treturn query\n")
	g.Printf("\t}\n")
	g.Printf("}\n\n")

	g.Printf("func %[1]sSort(sortFunc func(%[2]s []*%[1]s)) func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name, m.NameLowerCamelPlural)
	g.Printf("\treturn func(query *%[1]sQuery) *%[1]sQuery {\n", m.Name)
	g.Printf("\t\tif query != nil {\n")
	g.Printf("\t\t\tquery.sort = sortFunc\n")
	g.Printf("\t\t}\n")
	g.Printf("\t\treturn query\n")
	g.Printf("\t}\n")
	g.Printf("}\n\n")

	g.Printf("func Get%[2]s(opts ...%[1]sQueryOption) ([]*%[1]s, error) {\n", m.Name, m.NamePlural)
	g.Printf("\t%[2]sQuery := &%[1]sQuery{}\n", m.Name, m.NameLowerCamel)
	g.Printf("\tfor _, opt := range opts {\n")
	g.Printf("\t\t%[1]sQuery = opt(%[1]sQuery)\n", m.NameLowerCamel)
	g.Printf("\t}\n")
	g.Printf("\t_%s_mutex.RLock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.RUnlock()\n", m.Name)
	g.Printf("\t%[2]s := []*%[1]s{}\n", m.Name, m.NameLowerCamelPlural)
	g.Printf("\tif %sQuery.filter != nil {\n", m.NameLowerCamel)
	g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
	g.Printf("\t\t\tif %sQuery.filter(v) {\n", m.NameLowerCamel)
	g.Printf("\t\t\t\t%[1]s = append(%[1]s, v)\n", m.NameLowerCamelPlural)
	g.Printf("\t\t\t}\n")
	g.Printf("\t\t}\n")
	g.Printf("\t} else {\n")
	g.Printf("\t\tfor _, v := range _%s_cache {\n", m.Name)
	g.Printf("\t\t\t%[1]s = append(%[1]s, v)\n", m.NameLowerCamelPlural)
	g.Printf("\t\t}\n")
	g.Printf("\t}\n")
	g.Printf("\tif %sQuery.sort != nil {\n", m.NameLowerCamel)
	g.Printf("\t\t%sQuery.sort(%s)\n", m.NameLowerCamel, m.NameLowerCamelPlural)
	g.Printf("\t}\n")
	g.Printf("\treturn %s, nil\n", m.NameLowerCamelPlural)
	g.Printf("}\n\n")
}

func (g *Generator) outputAdd(m model, o option) {
	pkNamesWithModelPrefix := []string{}
	nonKeyFields := []string{}
	for _, f := range m.Fields {
		if f.PrimaryKey {
			pkNamesWithModelPrefix = append(pkNamesWithModelPrefix, m.NameLowerCamel+"."+f.Name)
		} else {
			nonKeyFields = append(nonKeyFields, f.NameLowerCamel+" "+f.Typ)
		}
	}

	g.Printf("func Add%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, strings.Join(nonKeyFields, ", "))
	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, f := range m.Fields {
		if !f.PrimaryKey && f.Typ == "string" {
			if f.Unique {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s, nil); err != nil {\n", m.Name, f.Name, f.NameLowerCamel)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")	
			} else {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLowerCamel)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")	
			}
		}
	}

	g.Printf("\t%[2]s := &%[1]s{\n", m.Name, m.NameLowerCamel)

	// TODO: child model
	for _, f := range m.Fields {
		if f.PrimaryKey {
			g.Printf("\t\t%[2]s: _%[1]s_maxRowNo + %[3]d,\n", m.Name, f.Name, o.Initial)
		} else {
			g.Printf("\t\t%s: %s,\n", f.Name, f.NameLowerCamel)
		}
	}

	g.Printf("\t}\n")
	g.Printf("\tif err := %s._asyncUpdate(); err != nil {\n", m.NameLowerCamel)
	g.Printf("\t\treturn nil, err\n")
	g.Printf("\t}\n")
	g.Printf("\t_%s_maxRowNo++\n", m.Name)
	g.Printf("\t_%[1]s_cache[%[3]s] = %[2]s\n", m.Name, m.NameLowerCamel, strings.Join(pkNamesWithModelPrefix, "]["))
	g.Printf("\t_%[1]s_rowNoMap[%[2]s] = _%[1]s_maxRowNo\n", m.Name, strings.Join(pkNamesWithModelPrefix, "]["))
	g.Printf("\treturn %s, nil\n", m.NameLowerCamel)
	g.Printf("}\n\n")
}

func (g *Generator) outputUpdate(m model) {
	pkNames := []string{}
	fNameAndTypes := []string{}
	for _, f := range m.Fields {
		if f.PrimaryKey {
			pkNames = append(pkNames, f.NameLowerCamel)
		}
		fNameAndTypes = append(fNameAndTypes, f.NameLowerCamel+" "+f.Typ)
	}

	g.Printf("func Update%[1]s(%[2]s) (*%[1]s, error) {\n", m.Name, strings.Join(fNameAndTypes, ", "))
	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, f := range m.Fields {
		if !f.PrimaryKey && f.Typ == "string" {
			if f.Unique {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s, &%[4]s); err != nil {\n", m.Name, f.Name, f.NameLowerCamel, strings.Join(pkNames, ", &"))
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")	
			} else {
				g.Printf("\tif err := _%[1]s_validate%[2]s(%[3]s); err != nil {\n", m.Name, f.Name, f.NameLowerCamel)
				g.Printf("\t\treturn nil, err\n")
				g.Printf("\t}\n")	
			}
		}
	}

	g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s]\n", m.Name, m.NameLowerCamel, strings.Join(pkNames, "]["))
	g.Printf("\tif !ok {\n")
	g.Printf("\t\treturn nil, &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("\t}\n")
	g.Printf("\t%[1]sCopy := *%[1]s\n", m.NameLowerCamel)

	for _, f := range m.Fields {
		if !f.PrimaryKey {
			g.Printf("\t%sCopy.%s = %s\n", m.NameLowerCamel, f.Name, f.NameLowerCamel)
		}
	}

	g.Printf("\tif err := (&%sCopy)._asyncUpdate(); err != nil {\n", m.NameLowerCamel)
	g.Printf("\t\treturn nil, err\n")
	g.Printf("\t}\n")
	g.Printf("\t%[1]s = &%[1]sCopy\n", m.NameLowerCamel)
	g.Printf("\treturn %s, nil\n", m.NameLowerCamel)
	g.Printf("}\n\n")
}

func (g *Generator) outputDelete(m model) {
	pkNames := []string{}
	pkNameAndTypes := []string{}
	for _, f := range m.PrimaryKeys() {
		pkNames = append(pkNames, f.NameLowerCamel)
		pkNameAndTypes = append(pkNameAndTypes, f.NameLowerCamel+" "+f.Typ)
	}
	childNames := []string{}
	for _, v := range m.Children {
		childNames = append(childNames, v.NameLowerCamelPlural)
	}

	g.Printf("func Delete%s(%s) error {\n", m.Name, strings.Join(pkNameAndTypes, ", "))
	g.Printf("\t_%s_mutex.Lock()\n", m.Name)
	g.Printf("\tdefer _%s_mutex.Unlock()\n", m.Name)

	for _, child := range m.Children {
		g.Printf("\t_%s_mutex.Lock()\n", child.Name)
		g.Printf("\tdefer _%s_mutex.Unlock()\n", child.Name)
	}

	g.Printf("\t%[2]s, ok := _%[1]s_cache[%[3]s]\n", m.Name, m.NameLowerCamel, strings.Join(pkNames, "]["))
	g.Printf("\tif !ok {\n")
	g.Printf("\t\treturn &sheetdb.NotFoundError{Model: \"%s\"}\n", m.Name)
	g.Printf("\t}\n")

	for _, child := range m.Children {
		g.Printf("\tvar %[2]s []*%[1]s\n", child.Name, child.NameLowerCamelPlural)
		g.Printf("\tfor _, v := range _%s_cache[%s] {\n", child.Name, strings.Join(pkNames, "]["))
		g.Printf("\t\t%[1]s = append(%[1]s, v)\n", child.NameLowerCamelPlural)
		g.Printf("\t}\n")	
	}

	g.Printf("\tif err := %s._asyncDelete(%s); err != nil {\n", m.NameLowerCamel, strings.Join(childNames, ", "))
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")

	if m.Parent == nil {
		g.Printf("\tdelete(_%[1]s_cache, %[2]s)\n", m.Name, pkNames[0])
	} else {
		// TODO
		g.Printf("\tdelete(_%[1]s_cache, %[2]s)\n", m.Name, pkNames[0])
	}

	for _, child := range m.Children {
		g.Printf("\tdelete(_%s_cache, %s)\n", child.Name, strings.Join(pkNames, ", "))
	}

	g.Printf("\treturn nil\n")
	g.Printf("}\n\n")
}

func (g *Generator) outputValidate(m model) {
}

func (g *Generator) outputParse(m model) {
}

func (g *Generator) outputAsync(m model, o option) {
}
