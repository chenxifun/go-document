package astgo

import (
	"github.com/chenxifun/go-document/file"
	"github.com/chenxifun/go-document/types"
	"github.com/chenxifun/logger"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
)

func ReadGoFile(goFile types.GoFile) (data *types.PkgData, err error) {
	data = &types.PkgData{
		FileInfo: goFile,
		PkgPath:  goFile.PkgPath,
	}
	f, err := parseFile(goFile.FilePath)
	if err != nil {
		logger.Error("ast parse file %s has error %v", goFile.FilePath, err)
		return nil, errors.WithMessagef(err, "ast parse file %s has error", goFile.FilePath)
	}

	data.PackName = f.Name.Name
	imports := ParseImports(f.Imports)
	data.Imports = imports

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			gd := parseGenDecl(d)
			if gd != nil {
				gd.SetPkgPath(goFile.PkgPath)
				data.Structs = append(data.Structs, gd)
			}

		case *ast.FuncDecl:
			fd := parseFuncDecl(d)
			if fd != nil && IsUp(fd.FuncName) {
				fd.SetPkgPath(goFile.PkgPath, imports)
				fd.BuildDoc()
				data.Funcs = append(data.Funcs, fd)
			}

		}

	}

	return
}

func parseFile(fileName string) (*ast.File, error) {
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

func IsUp(s string) bool {
	return s[0] >= 'A' && s[0] <= 'Z'
}
