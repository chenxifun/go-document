package go_document

import (
	"fmt"
	"github.com/chenxifun/go-document/astgo"
	"github.com/chenxifun/go-document/file"
	"github.com/chenxifun/go-document/types"
	"github.com/pkg/errors"
	"path"
	"strings"
)

type Doc struct {
	baseDir  string
	pkgPaths []string
	goFiles  []types.GoFile

	Packages map[string]*types.PkgData
	Structs  map[string]*types.StructData
	Funcs    map[string]*types.FuncData
}

func (d *Doc) SetBaseDir(dir string) *Doc {

	dir = strings.ReplaceAll(dir, "\\", "/")
	d.baseDir = dir
	return d
}

func (d *Doc) AddPkgPath(dir string) *Doc {
	dir = strings.ReplaceAll(dir, "\\", "/")
	d.pkgPaths = append(d.pkgPaths, dir)
	return d
}

func (d *Doc) AddPkgPaths(dirs []string) *Doc {
	for _, dir := range dirs {
		d.AddPkgPath(dir)
	}
	return d
}

func (d *Doc) readGoFile() {
	for _, pkgPath := range d.pkgPaths {
		dir := path.Join(d.baseDir, pkgPath)
		ns, err := file.ReadDir(dir, true)
		if err == nil {
			d.goFiles = append(d.goFiles, types.NewGoFiles(d.baseDir, ns)...)
		} else {
			fmt.Println(err)
		}
	}
}

func (d *Doc) Build() error {
	d.readGoFile()
	if d.Packages == nil {
		d.Packages = make(map[string]*types.PkgData)
	}
	for _, gf := range d.goFiles {
		if !gf.IsTestGo {
			data, err := astgo.ReadGoFile(gf)

			if err != nil {
				return errors.WithMessagef(err, "build go file %s has error", gf.FilePath)
			}
			pd, ok := d.Packages[gf.PkgPath]
			if ok {
				pd.Append(data)
			} else {
				d.Packages[gf.PkgPath] = data
			}
		}
	}

	return nil
}
