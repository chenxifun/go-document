package astgo

import (
	"fmt"
	"github.com/chenxifun/go-document/file"
	"github.com/chenxifun/go-document/types"
	"path"
	"testing"
)

func TestName(t *testing.T) {

	baseDir := "D:\\GoPath\\src\\"
	pkgPath := "github.com/chenxifun/jsonrpc/test"
	ns, err := file.ReadDir(path.Join(baseDir, pkgPath), false)

	if err != nil {
		t.Fatal(err)
	}

	for _, n := range ns {
		if !file.IsGoTestFile(n) {
			gof, _ := types.NewGoFile(baseDir, n)

			data, err := ReadGoFile(gof)
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println(data.PackName)
		}
		//fmt.Println(n,file.IsGoFile(n),file.IsGoTestFile(n))
	}
}
