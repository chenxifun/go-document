package types

type PkgData struct {
	FileInfo GoFile
	PackName string
	PkgPath  string
	Imports  map[string]ImportData
	Structs  []*StructData
	Funcs    []*FuncData
}

func (p *PkgData) Append(d *PkgData) {
	p.Structs = append(p.Structs, d.Structs...)
	p.Funcs = append(p.Funcs, d.Funcs...)
}

type PackageData struct {
	PackName string
	PkgPath  string
	Structs  map[string]*StructData
	Funcs    map[string]*FuncData
}

func (p *PackageData) Append(d *PkgData) {

}
