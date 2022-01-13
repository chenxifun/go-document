package go_document

import (
	"fmt"
	"testing"
)

func TestReadGofile(t *testing.T) {
	doc := &Doc{}

	doc.SetBaseDir("D:\\GoPath\\src\\").AddPkgPath("github.com\\chenxifun\\jsonrpc\\test").AddPkgPath("github.com\\chenxifun\\jsonrpc\\rpc")

	doc.Build()

	for _, g := range doc.goFiles {
		fmt.Println(g)
	}

}
