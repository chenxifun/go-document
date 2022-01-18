package types

import (
	"fmt"
	"github.com/chenxifun/go-document/doc"
)

type FuncData struct {
	Doc         string
	Title       string
	Description string
	Return      string
	FuncName    string
	StructName  string
	PkgPath     string
	Params      []*DeclField
	Results     []*DeclField
}

func (s *FuncData) SetPkgPath(pkgPath string, imports map[string]ImportData) {
	s.PkgPath = pkgPath
	for i, _ := range s.Params {
		setAliasPkg(s.Params[i], pkgPath, imports)
	}

	for i, _ := range s.Results {
		setAliasPkg(s.Results[i], pkgPath, imports)
	}
}

func (s *FuncData) ID() string {
	if s.StructName == "" {
		return fmt.Sprintf("%s.%s", s.PkgPath, s.FuncName)
	} else {
		return fmt.Sprintf("%s.%s.%s", s.PkgPath, s.StructName, s.FuncName)
	}

}

func (f *FuncData) BuildDoc() {
	if f.Doc == "" {
		return
	}
	d := doc.BuildDoc(f.Doc)

	f.Title = d.Title
	f.Description = d.Description
	f.Return = d.Returns

	for i, _ := range f.Params {
		f.Params[i].Doc = d.Params[f.Params[i].Name]
	}

}

func setAliasPkg(fs *DeclField, thisPkg string, imports map[string]ImportData) {

	alia := fs.Field.GetAlias()
	if alia == "" && fs.Field.IsObj() { //
		fs.Field.SetPkgPath(thisPkg)
	}

	if alia != "" {
		id, ok := imports[alia]
		if ok {
			fs.Field.SetPkgPath(id.Path)
		}
	}

}
