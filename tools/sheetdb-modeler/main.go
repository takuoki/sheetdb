package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeName      = flag.String("type", "", "type name; must be set")
	parentName    = flag.String("parent", "", "parent type name")
	childrenNames = flag.String("children", "", "comma-separated list of children type names")
	clientName    = flag.String("client", "dbClient", "variable name of sheetdb package client")
	modelSetName  = flag.String("modelset", "default", "model set name")
	initialNum    = flag.Int("initial", 1, "initial number of auto numbering")
	output        = flag.String("output", "", "output file name; default srcdir/<type>_model.go")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of sheetdb-modeler:\n")
	fmt.Fprintf(os.Stderr, "\tsheetdb-modeler [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tsheetdb-modeler [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttps://github.com/takuoki/sheetdb/tools/sheetdb-modeler\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("sheetdb-modeler: ")
	flag.Usage = usage
	flag.Parse()
	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		log.Fatal("argument must be one, and must be an existing directory name.")
	}
	g := generator{}
	g.parsePackage(dir)

	// Run generate.
	g.generate(*typeName, *parentName, *childrenNames, *clientName, *modelSetName, *initialNum)

	// Format the output.
	src := g.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("model_%s.go", *typeName)
		outputName = filepath.Join(dir, strings.ToLower(baseName))
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

type generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *pkg         // Package we are scanning.
}

func (g *generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

type file struct {
	pkg  *pkg      // Package to which this file belongs.
	file *ast.File // Parsed AST.
}

type pkg struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*file
}

func (g *generator) parsePackage(dir string) {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
	}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

func (g *generator) addPackage(p *packages.Package) {
	g.pkg = &pkg{
		name:  p.Name,
		defs:  p.TypesInfo.Defs,
		files: make([]*file, len(p.Syntax)),
	}

	for i, f := range p.Syntax {
		g.pkg.files[i] = &file{
			file: f,
			pkg:  g.pkg,
		}
	}
}

// format returns the gofmt-ed contents of the generator's buffer.
func (g *generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
