package types

import (
	"github.com/chenxifun/go-document/file"
	"strings"
)

type GoFile struct {
	FileName string
	FilePath string
	PkgPath  string
	IsTestGo bool
}

func NewGoFiles(baseDir string, filePaths []string) (goFile []GoFile) {

	for _, f := range filePaths {
		gof, isGo := NewGoFile(baseDir, f)
		if isGo {
			goFile = append(goFile, gof)
		}
	}
	return
}

func NewGoFile(baseDir, filePath string) (goFile GoFile, isGo bool) {

	baseDir = strings.ReplaceAll(baseDir, "\\", "/")
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	if len(baseDir) <= 0 {
		return
	}

	if baseDir[len(baseDir)-1:] != "/" {
		baseDir = baseDir + "/"
	}
	isGo = file.IsGoFile(filePath)
	if !isGo {
		return
	}
	goFile.FilePath = filePath

	last := strings.LastIndex(filePath, "/")
	goFile.FileName = filePath[last+1:]
	goFile.PkgPath = filePath[len(baseDir):last]
	goFile.IsTestGo = file.IsGoTestFile(filePath)

	return
}
