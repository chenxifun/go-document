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

func (p *PkgData) FindStruct(name string) *StructData {
	for i, s := range p.Structs {
		if s.Name == name {
			return p.Structs[i]
		}
	}
	return nil
}

func (p *PkgData) FindFunc(sname, fname string) *FuncData {
	for i, s := range p.Funcs {
		if s.StructName == sname && s.FuncName == fname {
			return p.Funcs[i]
		}
	}
	return nil
}
