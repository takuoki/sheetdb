package main

import (
	"go/ast"
	"log"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"github.com/takuoki/gocase"
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
	DirectChildrenNames      []string
}

type option struct {
	ClientName   string
	ModelSetName string
	Initial      int
}

type field struct {
	Name          string // UserID | Date
	NameLower     string // userID | date
	Typ           string // int    | *sheetdb.Date
	TypNonPointer string // int    | sheetdb.Date
	Package       string //        | sheetdb
	TypRaw        string // int    | Date
	IsPointer     bool
	IsPk          bool
	IsParentKey   bool
	AllowEmpty    bool
	Unique        bool
}

func (g *Generator) generate(typeName, parentName, childrenNames, clientName, modelSetName string, initialNum int) {
	s := search{
		Typ:    typeName,
		Parent: parentName,
	}
	for _, c := range strings.Split(childrenNames, ",") {
		if c != "" {
			s.Children = append(s.Children, c)
		}
	}
	for _, f := range g.pkg.files {
		if f.file != nil {
			ast.Inspect(f.file, s.searchType)
			ast.Inspect(f.file, s.searchParent)
			ast.Inspect(f.file, s.searchChildren)
		}
	}
	if s.Model == nil {
		log.Fatalf("unable to find specified type. (type=%s)", typeName)
	}
	if parentName != "" && s.ParentModel == nil {
		log.Fatalf("unable to find specified parent type. (type=%s)", parentName)
	}
	if len(s.Children) != len(s.ChildrenModels) {
		log.Fatal("unable to find some specified children types.")
	}

	s.Model.Parent = s.ParentModel
	s.Model.Children = s.ChildrenModels
	s.Model.ChildrenNames = s.Children
	for _, c := range s.Children {
		s.Model.ChildrenNamePlurals = append(s.Model.ChildrenNamePlurals, inflection.Plural(c))
		s.Model.ChildrenNameLowers = append(s.Model.ChildrenNameLowers, gocase.To(strcase.ToLowerCamel(gocase.Revert(c))))
		s.Model.ChildrenNameLowerPlurals = append(s.Model.ChildrenNameLowerPlurals, gocase.To(inflection.Plural(strcase.ToLowerCamel(gocase.Revert(c)))))
	}

	opt := option{
		ClientName:   clientName,
		ModelSetName: modelSetName,
		Initial:      initialNum,
	}

	if err := g.validate(*s.Model, opt); err != nil {
		log.Fatal(err)
	}

	g.output(*s.Model, opt)
}

type search struct {
	Typ      string
	Parent   string
	Children []string

	Model          *model
	ParentModel    *model
	ChildrenModels []model
}

func (s *search) searchType(node ast.Node) bool {
	typ, ok := node.(*ast.TypeSpec)
	if !ok || typ.Name.Name != s.Typ {
		return true
	}
	s.Model = s.buildModel(typ)
	return false
}

func (s *search) searchParent(node ast.Node) bool {
	typ, ok := node.(*ast.TypeSpec)
	if !ok || typ.Name.Name != s.Parent {
		return true
	}
	s.ParentModel = s.buildModel(typ)
	return false
}

func (s *search) searchChildren(node ast.Node) bool {
	typ, ok := node.(*ast.TypeSpec)
	if !ok {
		return true
	}
	ok = false
	for _, c := range s.Children {
		if typ.Name.Name == c {
			ok = true
			break
		}
	}
	if !ok {
		return true
	}
	s.ChildrenModels = append(s.ChildrenModels, *(s.buildModel(typ)))
	return false
}

func (s *search) buildModel(typ *ast.TypeSpec) *model {

	m := model{
		Name:            typ.Name.Name,
		NamePlural:      inflection.Plural(typ.Name.Name),
		NameLower:       gocase.To(strcase.ToLowerCamel(gocase.Revert(typ.Name.Name))),
		NameLowerPlural: gocase.To(inflection.Plural(strcase.ToLowerCamel(gocase.Revert(typ.Name.Name)))),
	}

	st, ok := typ.Type.(*ast.StructType)
	if !ok {
		log.Fatalf("specified type is not struct type. (type=%s)", typ.Name.Name)
	}

	for _, f := range st.Fields.List {
		if len(f.Names) != 1 {
			log.Fatalf("specify one field per line. (type=%s, fields=%v)", typ.Name.Name, f.Names)
		}
		f2 := field{
			Name:      f.Names[0].Name,
			NameLower: gocase.To(strcase.ToLowerCamel(gocase.Revert(f.Names[0].Name))),
		}
		m.FieldNames = append(m.FieldNames, f2.Name)
		m.FieldNameLowers = append(m.FieldNameLowers, f2.NameLower)

		// set type related fields
		switch ft := f.Type.(type) {
		case *ast.Ident:
			f2.Typ = ft.Name
			f2.TypNonPointer = ft.Name
			f2.TypRaw = ft.Name
		case *ast.SelectorExpr:
			x, ok := ft.X.(*ast.Ident)
			if !ok {
				log.Fatalf("package name is invalid. (type=%s, field=%s)", typ.Name.Name, f.Names[0].Name)
			}
			f2.Typ = x.Name + "." + ft.Sel.Name
			f2.TypNonPointer = x.Name + "." + ft.Sel.Name
			f2.Package = x.Name
			f2.TypRaw = ft.Sel.Name
		case *ast.StarExpr:
			f2.IsPointer = true
			switch x := ft.X.(type) {
			case *ast.Ident:
				f2.Typ = "*" + x.Name
				f2.TypNonPointer = x.Name
				f2.TypRaw = x.Name
			case *ast.SelectorExpr:
				x2, ok := x.X.(*ast.Ident)
				if !ok {
					log.Fatalf("package name is invalid. (type=%s, field=%s)", typ.Name.Name, f.Names[0].Name)
				}
				f2.Typ = "*" + x2.Name + "." + x.Sel.Name
				f2.TypNonPointer = x2.Name + "." + x.Sel.Name
				f2.Package = x2.Name
				f2.TypRaw = x.Sel.Name
			default:
				log.Fatalf("This type is unsupported (model=%s, field=%s)", typ.Name.Name, f.Names[0].Name)
			}
		default:
			log.Fatalf("This type is unsupported (model=%s, field=%s)", typ.Name.Name, f.Names[0].Name)
		}
		m.FieldTypes = append(m.FieldTypes, f2.Typ)

		// set tag related fields
		if f.Tag != nil && f.Tag.Value != "" {
			for _, tags := range strings.Split(f.Tag.Value[1:len(f.Tag.Value)-1], " ") {
				if len(tags) < 4 || tags[:4] != `db:"` {
					continue
				}
				for _, tag := range strings.Split(tags[4:len(tags)-1], ",") {
					switch tag {
					case "primarykey":
						f2.IsPk = true
						m.PkNames = append(m.PkNames, f2.Name)
						m.PkNameLowers = append(m.PkNameLowers, f2.NameLower)
						m.PkTypes = append(m.PkTypes, f2.Typ)
						// the last primary key is "this" key
						m.ThisKeyName = f2.Name
						m.ThisKeyNameLower = f2.NameLower
						m.ThisKeyType = f2.Typ
					case "allowempty":
						f2.AllowEmpty = true
					case "unique":
						f2.Unique = true
					}
				}
				break
			}
		}
		if !f2.IsPk {
			m.NonPkNameLowers = append(m.NonPkNameLowers, f2.NameLower)
			m.NonPkTypes = append(m.NonPkTypes, f2.Typ)
		}

		m.Fields = append(m.Fields, f2)
	}
	for i := 0; i < len(m.Fields); i++ {
		if m.Fields[i].Name != m.ThisKeyName {
			m.Fields[i].IsParentKey = true
		}
	}

	return &m
}
