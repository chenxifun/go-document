package go_docs

import (
	"fmt"
	"github.com/chenxifun/go-document/astgo"
	"github.com/chenxifun/go-document/file"
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(fileName string) (*ast.File, error) {
	src, err := file.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ParseAst(pkgPath string, f *ast.File) (err error) {

	packageName := f.Name.Name
	fmt.Println(packageName)

	imports := astgo.ParseImports(f.Imports)
	fmt.Println(imports)

	for _, d := range f.Decls {

		ast.FilterDecl(d, func(s string) bool {

			return s == ""
		})

	}

	return

}

func parseDecl(decl ast.Decl) (err error) {

	switch d := decl.(type) {
	case *ast.GenDecl:
		paseGenDecl(d)
	case *ast.FuncDecl:
		pasefuncDecl(d)
	}

	return
}

func paseGenDecl(decl *ast.GenDecl) {

}

func pasefuncDecl(decl *ast.FuncDecl) {

}
