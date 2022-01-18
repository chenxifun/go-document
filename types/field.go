package types

type Fielder interface {
	SetPkgPath(pkgPath string)
	GetPkgPath() string
	SetPkgName(pkgName string)
	GetPkgName() string
	SetAlias(alias string)
	GetAlias() string
	IsObj() bool
	ID() string
	Type() string
}

type FieldType struct {
	// 别名
	Alias string
	//包路径
	PkgPath string
	//类型名
	PkgName string
	// 是否 struct
	Obj bool
}

func (f *FieldType) Type() string {
	return "Obj"
}

func (f *FieldType) ID() string {

	var p string
	if f.PkgPath == "" && f.Alias != "" {
		p = f.Alias
	} else {
		p = f.PkgPath
	}

	if p == "" {
		return f.PkgName
	}
	return p + "." + f.PkgName
}

func (f *FieldType) IsObj() bool {
	return f.Obj
}

func (f *FieldType) SetAlias(alias string) {
	f.Alias = alias
}
func (f *FieldType) GetAlias() string {
	return f.Alias
}
func (f *FieldType) SetPkgPath(pkgPath string) {
	f.PkgPath = pkgPath
}

func (f *FieldType) GetPkgPath() string {
	return f.PkgPath
}

func (f *FieldType) SetPkgName(pkgName string) {
	f.PkgName = pkgName
}

func (f *FieldType) GetPkgName() string {
	return f.PkgName
}

type ArrayType struct {
	Field Fielder
	Len   int
}

func (f *ArrayType) Type() string {
	return "Array"
}
func (f *ArrayType) ID() string {
	return f.Field.ID()
}

func (f *ArrayType) IsObj() bool {
	return f.Field.IsObj()
}
func (f *ArrayType) SetAlias(alias string) {
	f.Field.SetAlias(alias)
}
func (f *ArrayType) GetAlias() string {
	return f.Field.GetAlias()
}
func (f *ArrayType) SetPkgPath(pkgPath string) {
	f.Field.SetPkgPath(pkgPath)
}

func (f *ArrayType) GetPkgPath() string {
	return f.Field.GetPkgPath()
}

func (f *ArrayType) SetPkgName(pkgName string) {
	f.Field.SetPkgName(pkgName)
}

func (f *ArrayType) GetPkgName() string {
	return f.Field.GetPkgPath()
}

type MapType struct {
	Key   Fielder
	Value Fielder
}

func (f *MapType) Type() string {
	return "Map"
}

func (f *MapType) ID() string {
	return f.Value.ID()
}

func (f *MapType) IsObj() bool {
	return f.Value.IsObj()
}

func (f *MapType) SetAlias(alias string) {
	f.Value.SetAlias(alias)
}
func (f *MapType) GetAlias() string {
	return f.Value.GetAlias()
}
func (f *MapType) SetPkgPath(pkgPath string) {
	f.Value.SetPkgPath(pkgPath)
}

func (f *MapType) GetPkgPath() string {
	return f.Value.GetPkgPath()
}

func (f *MapType) SetPkgName(pkgName string) {
	f.Value.SetPkgName(pkgName)
}

func (f *MapType) GetPkgName() string {
	return f.Value.GetPkgPath()
}
