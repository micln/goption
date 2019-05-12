package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/Kretech/xgo/astutil"
	"github.com/pkg/errors"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Package struct {
	*doc.Package
}

type ClassField struct {
	ast.Field
}

func (f *ClassField) Name() string {
	return f.Names[0].String()
}

func (f *ClassField) TypeName() string {
	return astutil.ExprString(f.Type)
}

type Class struct {
	ast.StructType

	Package Package

	Name   string
	Fields []ClassField
}

func getPackageClasses(path string) (classes []*Class, err error) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, func(fileInfo os.FileInfo) bool {
		return strings.HasSuffix(fileInfo.Name(), ".go") ||
			!strings.HasSuffix(fileInfo.Name(), "_gen.go") ||
			!strings.HasSuffix(fileInfo.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		err = errors.Wrap(err, `parser.ParseDir`)
		return
	}

	for _, astPkg := range pkgs {
		docPkg := doc.New(astPkg, path, doc.AllDecls)
		pkg := Package{docPkg}
		for _, ptyp := range docPkg.Types {
			for _, spec := range ptyp.Decl.Specs {
				typeSpec := spec.(*ast.TypeSpec)

				structTypePtr, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				newClass := &Class{
					Package: pkg,
					Name:    typeSpec.Name.String(),
				}
				newClass.StructType = *structTypePtr
				for _, filed := range newClass.StructType.Fields.List {
					newField := *filed
					newClass.Fields = append(newClass.Fields, ClassField{newField})
				}
				classes = append(classes, newClass)
			}
		}
	}

	return
}

func parsePackageOld(path string) (pkgName string, files []*ast.File, err error) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), path, func(fileInfo os.FileInfo) bool {
		return strings.HasSuffix(fileInfo.Name(), ".go")
	}, parser.ParseComments)

	for name, pkg := range pkgs {
		pkgName = name

		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}

	return
}

func parseClass(file *ast.File) (classes []*Class, err error) {
	ast.Inspect(file, func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok /*|| decl.Tok != token.TYPE*/ {
			// We only care about const declarations.
			return true
		}

		for _, spec := range decl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			newClass := &Class{Name: typeSpec.Name.String()}
			newClass.StructType = *typeSpec.Type.(*ast.StructType)
			for _, filed := range newClass.StructType.Fields.List {
				newField := *filed
				newClass.Fields = append(newClass.Fields, ClassField{newField})
			}
			classes = append(classes, newClass)
		}

		return false
	})
	return
}
