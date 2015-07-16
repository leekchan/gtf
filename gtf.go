package gtf

import (
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"strings"
)

// recovery will silently swallow all unexpected panics.
func recovery() {
	recover()
}

var GtfFuncMap = template.FuncMap{
	"stringReplace": func(s1 string, s2 string) string {
		defer recovery()

		return strings.Replace(s2, s1, "", -1)
	},
	"stringDefault": func(s1 string, s2 string) string {
		defer recovery()

		if len(s2) > 0 {
			return s2
		}
		return s1
	},
	"stringLength": func(s string) int {
		defer recovery()

		return len(s)
	},
	"stringLower": func(s string) string {
		defer recovery()

		return strings.ToLower(s)
	},
	"stringUpper": func(s string) string {
		defer recovery()

		return strings.ToUpper(s)
	},
	"stringTruncatechars": func(n int, s string) string {
		defer recovery()

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
	"stringUrlencode": func(s string) string {
		defer recovery()

		return url.QueryEscape(s)
	},
	"stringWordcount": func(s string) int {
		defer recovery()

		return len(strings.Fields(s))
	},
	"intDivisibleby": func(i1 int, i2 int) bool {
		defer recovery()

		return i2%i1 == 0
	},
	"stringLengthIs": func(i int, s string) bool {
		defer recovery()

		return i == len([]rune(s))
	},
	"stringTrim": func(s string) string {
		defer recovery()

		return strings.TrimSpace(s)
	},
	"stringCapfirst": func(s string) string {
		defer recovery()

		return strings.ToUpper(string(s[0])) + s[1:]
	},
	"intPluralize": func(arg string, value int) string {
		defer recovery()

		if !strings.Contains(arg, ",") {
			arg = "," + arg
		}

		bits := strings.Split(arg, ",")

		if len(bits) > 2 {
			return ""
		}

		if value == 1 {
			return bits[0]
		}

		return bits[1]
	},
	"boolYesno": func(yes string, no string, value bool) string {
		defer recovery()

		if value {
			return yes
		}

		return no
	},
	"stringRjust": func(arg int, value string) string {
		defer recovery()

		n := arg - len([]rune(value))

		if n > 0 {
			value = strings.Repeat(" ", n) + value
		}

		return value
	},
	"stringLjust": func(arg int, value string) string {
		defer recovery()

		n := arg - len([]rune(value))

		if n > 0 {
			value = value + strings.Repeat(" ", n)
		}

		return value
	},
	"stringCenter": func(arg int, value string) string {
		defer recovery()

		n := arg - len([]rune(value))

		if n > 0 {
			left := n / 2
			right := n - left
			value = strings.Repeat(" ", left) + value + strings.Repeat(" ", right)
		}

		return value
	},
	"filesizeformat": func(value interface{}) string {
		defer recovery()

		var size float64

		switch value.(type) {
		case int, int8, int16, int32, int64:
			size = float64(reflect.ValueOf(value).Int())
		case uint, uint8, uint16, uint32, uint64:
			size = float64(reflect.ValueOf(value).Uint())
		case float32, float64:
			size = reflect.ValueOf(value).Float()
		default:
			return ""
		}

		var KB float64 = 1 << 10
		var MB float64 = 1 << 20
		var GB float64 = 1 << 30
		var TB float64 = 1 << 40
		var PB float64 = 1 << 50

		filesizeFormat := func(filesize float64, suffix string) string {
			return strings.Replace(fmt.Sprintf("%.1f %s", filesize, suffix), ".0", "", -1)
		}

		var result string
		if size < KB {
			result = filesizeFormat(size, "bytes")
		} else if size < MB {
			result = filesizeFormat(size/KB, "KB")
		} else if size < GB {
			result = filesizeFormat(size/MB, "MB")
		} else if size < TB {
			result = filesizeFormat(size/GB, "GB")
		} else if size < PB {
			result = filesizeFormat(size/TB, "TB")
		} else {
			result = filesizeFormat(size/PB, "PB")
		}

		return result
	},
}

// gtf.New is a wrapper function of template.New(http://golang.org/pkg/text/template/#New).
// It automatically adds the gtf functions to the template's function map
// and returns template.Template(http://golang.org/pkg/text/template/#Template).
func New(name string) *template.Template {
	return template.New(name).Funcs(GtfFuncMap)
}
