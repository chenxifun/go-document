package doc

import (
	"fmt"
	"strings"
)

var keys = []string{"@title", "@description", "@param", "@return"}

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

	//for i,l :=range lines {
	//	l = strings.TrimSpace(l)
	//	if i !=0 {
	//		if !hasKey(l) && len(lines)>i-1 {
	//			lines[i-1] = fmt.Sprintf("%s\n%s",lines[i-1],l)
	//			lines = append(lines[:i],lines[i+1:] ...)
	//		}
	//	}
	//}

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
		}
	}

	return d
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
}
