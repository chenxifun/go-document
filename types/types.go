package types

import "fmt"

type StructType string

func (s StructType) String() string {
	return string(s)
}

var Struct_Type StructType = "struct"
var Interface_Type StructType = "interface"

type StructData struct {
	PkgPath string
	Name    string
	Doc     string
	Fields  []*DeclField
}

func (s *StructData) SetPkgPath(pkgPath string) {
	s.PkgPath = pkgPath
	for i, f := range s.Fields {
		if f.Field.GetAlias() == "" && f.Field.IsObj() {
			s.Fields[i].Field.SetPkgPath(pkgPath)
		}

	}
}

func (s *StructData) ID() string {
	return fmt.Sprintf("%s.%s", s.PkgPath, s.Name)
}

type DeclField struct {
	Doc  string
	Name string
	Tag  string

	Field Fielder
}
