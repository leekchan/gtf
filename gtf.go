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
}

func New(name string) *template.Template {
	return template.New(name).Funcs(GtfFuncMap)
}