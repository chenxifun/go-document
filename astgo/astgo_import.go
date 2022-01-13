package astgo

import (
	"github.com/chenxifun/go-document/types"
	"go/ast"
	"strings"
)

func ParseImports(imps []*ast.ImportSpec) map[string]types.ImportData {

	impData := make(map[string]types.ImportData)

	for _, s := range imps {
		d := parseImport(s)
		impData[d.Alias] = d

	}
	return impData
}

func parseImport(s *ast.ImportSpec) types.ImportData {

	d := types.ImportData{

		Path: strings.Trim(s.Path.Value, "\""),
	}

	if s.Name != nil {
		d.Alias = s.Name.Name
	}

	index := strings.LastIndex(d.Path, "/")
	if index == -1 {
		d.Package = d.Path
	} else {
		d.Package = d.Path[index+1:]
	}

	if strings.TrimSpace(d.Alias) == "" {
		d.Alias = d.Package
	}

	return d
}
