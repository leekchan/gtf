package gtf

import (
	"html/template"
    "strings"
)

var GtfFuncMap = template.FuncMap {
	"stringReplace": func(s1 string, s2 string) string {
		return strings.Replace(s2, s1, "", -1)
	},
	"stringDefault": func(s1 string, s2 string) string {
		if len(s2) > 0 {
			return s2
		}
		return s1
	},
	"stringLength": func(s string) int {
		return len(s)
	},
	"stringLower": func(s string) string {
		return strings.ToLower(s)
	},
	"stringTruncateChars": func(n int, s string) string {
		if n < 0 {
			return s
		}
		
		r := []rune(s)
		rLength := len(r)
		
		if n >= rLength {
			return s
		}
		
		if n > 3 && rLength > 3 {
			return string(r[:n-3]) + "..."
		}
		
		return string(r[:n])
	},
}

// gtf.New is a wrapper function of template.New(http://golang.org/pkg/text/template/#New). 
// It automatically adds the gtf functions to the template's function map 
// and returns template.Template(http://golang.org/pkg/text/template/#Template).
func New(name string) *template.Template {
	return template.New(name).Funcs(GtfFuncMap)
}