package main

import (
	"fmt"
)

func (g *generator) validate(m model, o option) error {
	if err := g.validatePrimaryKey(m); err != nil {
		return err
	}
	if err := g.validateChildren(m); err != nil {
		return err
	}
	for _, f := range m.Fields {
		if err := g.validateField(f, m); err != nil {
			return err
		}
	}
	if err := g.validateOption(m, o); err != nil {
		return err
	}
	return nil
}

func (g *generator) validatePrimaryKey(m model) error {
	if m.ThisKeyName == "" || m.ThisKeyType == "" {
		return fmt.Errorf("There are no primary key (model=%s)", m.Name)
	}
	for _, f := range m.Fields {
		if f.Name == m.ThisKeyName {
			if f.IsPointer {
				return fmt.Errorf("The type of primary key must not pointer (model=%s, field=%s)", m.Name, m.ThisKeyName)
			}
			if f.AllowEmpty {
				return fmt.Errorf("Must not put an 'allowempty' tag on primary key (model=%s, field=%s)", m.Name, m.ThisKeyName)
			}
			if f.Unique {
				return fmt.Errorf("Must not put an 'unique' tag on primary key (model=%s, field=%s)", m.Name, m.ThisKeyName)
			}
			break
		}
	}
	if m.Parent == nil {
		if len(m.PkNames) != 1 {
			return fmt.Errorf("If model doesn't have any parent, the number of primary key must be one (model=%s)", m.Name)
		}
		return nil
	}
	if len(m.PkNames) != len(m.Parent.PkNames)+1 {
		return fmt.Errorf("The number of primary key must be the number of primary key of parent + 1 (model=%s)", m.Name)
	}
	for i := 0; i < len(m.Parent.PkNames); i++ {
		if m.PkNames[i] != m.Parent.PkNames[i] {
			return fmt.Errorf("The name and order of primary key must be same as the parent one (model=%s, field=%s)", m.Name, m.PkNames[i])
		}
		if m.PkTypes[i] != m.Parent.PkTypes[i] {
			return fmt.Errorf("The type of primary key must be same as the parent one (model=%s, field=%s)", m.Name, m.PkNames[i])
		}
	}
	return nil
}

func (g *generator) validateChildren(m model) error {
	for _, child := range m.Children {
		if len(child.PkNames) <= len(m.PkNames) {
			return fmt.Errorf("The number of primary key must be greater than the number of primary key of parent (model=%s)", child.Name)
		}
		for i := 0; i < len(m.PkNames); i++ {
			if child.PkNames[i] != m.PkNames[i] {
				return fmt.Errorf("The name and order of primary key must be same as the parent one (model=%s, field=%s)", child.Name, child.PkNames[i])
			}
			if child.PkTypes[i] != m.PkTypes[i] {
				return fmt.Errorf("The type of primary key must be same as the parent one (model=%s, field=%s)", child.Name, child.PkNames[i])
			}
		}
	}
	return nil
}

func (g *generator) validateField(f field, m model) error {

	if f.Package == "" {
		if f.Typ == "*string" {
			return fmt.Errorf("This type is unsupported (model=%s, field=%s, type=%s)", m.Name, f.Name, f.Typ)
		}
		switch f.TypRaw {
		case "byte", "rune", "uintptr", "complex64", "complex128":
			return fmt.Errorf("This type is unsupported (model=%s, field=%s, type=%s)", m.Name, f.Name, f.Typ)
		}
	} else {
		if f.Package != "sheetdb" || (f.TypRaw != "Date" && f.TypRaw != "Datetime") {
			return fmt.Errorf("This type is unsupported (model=%s, field=%s, type=%s)", m.Name, f.Name, f.Typ)
		}
	}

	if f.AllowEmpty {
		if f.Typ != "string" {
			return fmt.Errorf("Must not put an 'allowempty' tag on non-string field (model=%s, field=%s)", m.Name, f.Name)
		}
	}

	if f.Unique {
		if f.Typ != "string" {
			return fmt.Errorf("Must not put an 'unique' tag on non-string field (model=%s, field=%s)", m.Name, f.Name)
		}
	}

	return nil
}

func (g *generator) validateOption(m model, o option) error {
	if o.ClientName == "" {
		return fmt.Errorf("Client name must not be empty string (model=%s)", m.Name)
	}
	if o.ModelSetName == "" {
		return fmt.Errorf("Model set name must not be empty string (model=%s)", m.Name)
	}
	if o.Initial <= 0 {
		return fmt.Errorf("Initial number must be positive (model=%s)", m.Name)
	}
	if !autoNumbering(m.ThisKeyName, m.ThisKeyType) && o.Initial != 1 { // default = 1
		return fmt.Errorf("Must not specify initial number option, because the key of this model is not for auto numbering (model=%s)", m.Name)
	}
	return nil
}
