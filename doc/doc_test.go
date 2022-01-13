package doc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBuildDoc(t *testing.T) {

	doc := `ssss
@title    函数名称
@description   函数的详细描述
@param ctx 参数注释
@param what 参数注释
 		换行
  aaa
aaa
@return 返回值注释
aaa
@sort 10
`

	ls := BuildDoc(doc)

	jb, _ := json.Marshal(&ls)

	fmt.Println(string(jb))
}

func TestBuildLine(t *testing.T) {

	k, c := splitLine("params ctx 参数注释")
	fmt.Println(k)
	fmt.Println(c)
}
