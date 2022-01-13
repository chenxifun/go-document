package doc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var keys = []string{"@title", "@description", "@param", "@return", "@sort"}
var defSort = 999

func BuildDoc(doc string) DocInfo {

	doc = strings.Trim(doc, "\n")
	doc = strings.ReplaceAll(doc, "\t", "")
	lines := strings.Split(doc, "\n")

	l := len(lines)
	i := 0

	for j := 1; j < l; j++ {
		if !hasKey(lines[j]) {
			lines[i] = fmt.Sprintf("%s\n%s", lines[i], lines[j])
			lines = append(lines[:i+1], lines[i+2:]...)
			j--
		} else {
			i++
		}

		if j >= len(lines)-1 {
			break
		}
	}
	return buildLines(lines)
}

func hasKey(line string) bool {
	return len(line) > 0 && line[0] == '@'
}

func buildLines(lines []string) DocInfo {
	d := DocInfo{
		Params: make(map[string]string),
	}

	for _, l := range lines {
		k, c := splitLine(l)
		switch k {
		case "title":
			d.Title = c
		case "description":
			d.Description = c
		case "param":
			p, doc := splitLine(c)
			d.Params[p] = doc
		case "return":
			d.Returns = c
		case "sort":
			d.Sort = sort(c)
		}
	}

	return d
}

func sort(s string) int {
	sort, err := strconv.Atoi(s)
	if err != nil {
		sort = defSort
	}
	return sort
}

func splitLine(line string) (key, context string) {
	line = strings.TrimSpace(line)
	spce := strings.Index(line, " ")
	if spce != -1 {
		if hasKey(line) {
			return line[1:spce], strings.TrimSpace(line[spce+1:])
		} else {
			return line[:spce], strings.TrimSpace(line[spce+1:])
		}
	}
	return line, ""
}

type DocInfo struct {
	Title       string
	Description string
	Params      map[string]string
	Returns     string
	Sort        int
}

func NewExampleData(kind reflect.Kind) string {

	switch kind {
	case reflect.Bool:
		return "true"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "123"
	case reflect.Float32, reflect.Float64:
		return "1.23"
	case reflect.String:
		return "\"abc\""
	default:
		return kind.String()
	}

}
