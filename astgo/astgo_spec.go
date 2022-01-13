package astgo

import (
	"github.com/chenxifun/go-document/types"
	"go/ast"
)

func parseGenDecl(d *ast.GenDecl) *types.StructData {

	if len(d.Specs) > 0 {
		sd := parseSpec(d.Specs[0])
		if sd != nil && d.Doc != nil {
			sd.Doc = d.Doc.Text()
		}
		return sd
	}
	return nil

}

func parseSpec(s ast.Spec) *types.StructData {
	ts, ok := s.(*ast.TypeSpec)
	if ok {
		return parseTypeSpec(ts)
	}

	return nil
}

func parseTypeSpec(t *ast.TypeSpec) *types.StructData {
	d := &types.StructData{
		Name: t.Name.Name,
	}
	switch st := t.Type.(type) {
	case *ast.StructType:
		d.Fields = parseStructType(st)
	case *ast.FuncType:
		// fun struct 丢弃
	}

	return d
}
