package astgo

import (
	"github.com/chenxifun/go-document/types"
	"go/ast"
)

func parseField(f *ast.Field, onlyPub bool) []*types.DeclField {

	var fds []*types.DeclField

	if len(f.Names) == 0 {

		sf := &types.DeclField{
			Field: parseExpr(f.Type),
		}

		if f.Tag != nil {
			sf.Tag = f.Tag.Value
		}

		if f.Doc != nil {
			sf.Doc = f.Doc.Text()
		}

		fds = append(fds, sf)
		return fds
	}

	// 一行有多个参数
	// 例如： a,b string
	for _, fi := range f.Names {
		k := fi.Name
		if !onlyPub || (k[0] >= 'A' && k[0] <= 'Z') {
			sf := &types.DeclField{
				Name:  k,
				Field: parseExpr(f.Type),
			}

			if f.Tag != nil {
				sf.Tag = f.Tag.Value
			}

			if f.Doc != nil {
				sf.Doc = f.Doc.Text()
			}

			fds = append(fds, sf)
		}
	}

	return fds
}

func parseExpr(e ast.Expr) types.Fielder {

	var ft types.Fielder
	if e == nil {
		return ft
	}

	switch d := e.(type) {
	case *ast.Ident:
		ft = parseIdent(d)
	case *ast.SelectorExpr:
		ft = parseSelectorExpr(d)
	case *ast.StarExpr:
		ft = parseStarExpr(d)
	case *ast.ArrayType:
		ft = parseArrayType(d)
	case *ast.MapType:
		ft = parseMapType(d)
	case *ast.ChanType:
		ft = parseChanType(d)
	default:
		ft = &types.FieldType{
			// 参数类型名
			PkgName: "",
		}
	}

	return ft
}

//参数类型
func parseIdent(i *ast.Ident) types.Fielder {
	d := &types.FieldType{
		// 参数类型名
		PkgName: i.Name,
	}

	if i.Obj != nil {
		d.Obj = true
		d.PkgName = i.Obj.Name
	}

	return d

}

func parseSelectorExpr(s *ast.SelectorExpr) types.Fielder {
	d := parseExpr(s.X)
	d.SetAlias(d.GetPkgName())
	if s.Sel != nil {
		d.SetPkgName(s.Sel.Name)
	}
	return d

}

func parseStarExpr(s *ast.StarExpr) types.Fielder {
	d := parseExpr(s.X)
	return d

}

func parseArrayType(a *ast.ArrayType) types.Fielder {
	d := &types.ArrayType{
		Field: parseExpr(a.Elt),
	}

	return d
}

func parseMapType(m *ast.MapType) types.Fielder {
	d := &types.MapType{

		Key:   parseExpr(m.Key),
		Value: parseExpr(m.Value),
	}

	return d
}

func parseChanType(c *ast.ChanType) types.Fielder {
	d := parseExpr(c.Value)
	return d
}

func parseStructType(s *ast.StructType) []*types.DeclField {

	return parseFieldList(s.Fields, true)

}

func parseFieldList(list *ast.FieldList, onlyPub bool) []*types.DeclField {
	var fs []*types.DeclField
	if list == nil {
		return fs
	}
	for _, f := range list.List {
		fs = append(fs, parseField(f, onlyPub)...)
	}

	return fs
}
