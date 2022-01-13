package astgo

import (
	"github.com/chenxifun/go-document/types"
	"go/ast"
)

func parseFuncDecl(f *ast.FuncDecl) *types.FuncData {
	fd := &types.FuncData{}
	fd.Doc = f.Doc.Text()
	fd.FuncName = f.Name.Name

	fd.Params = parseFieldList(f.Type.Params, false)
	fd.Results = parseFieldList(f.Type.Results, false)

	if f.Recv != nil && len(f.Recv.List) > 0 {

		s := parseExpr(f.Recv.List[0].Type)

		fd.StructName = s.GetPkgName()
	}

	return fd
}
