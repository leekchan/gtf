package gtf

import (
	"fmt"
	"html/template"
	"math"
	"net/url"
	"reflect"
	"strings"
)

// recovery will silently swallow all unexpected panics.
func recovery() {
	recover()
}

var GtfFuncMap = template.FuncMap{
	"replace": func(s1 string, s2 string) string {
		defer recovery()

		return strings.Replace(s2, s1, "", -1)
	},
	"default": func(arg interface{}, value interface{}) interface{} {
		defer recovery()

		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
			if v.Len() == 0 {
				return arg
			}
		case reflect.Bool:
			if !v.Bool() {
				return arg
			}
		default:
			return value
		}

		return value
	},
	"length": func(value interface{}) int {
		defer recovery()

		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			return v.Len()
		case reflect.String:
			return len([]rune(v.String()))
		}

		return 0
	},
	"lower": func(s string) string {
		defer recovery()

		return strings.ToLower(s)
	},
	"upper": func(s string) string {
		defer recovery()

		return strings.ToUpper(s)
	},
	"truncatechars": func(n int, s string) string {
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
	"urlencode": func(s string) string {
		defer recovery()

		return url.QueryEscape(s)
	},
	"wordcount": func(s string) int {
		defer recovery()

		return len(strings.Fields(s))
	},
	"divisibleby": func(arg interface{}, value interface{}) bool {
		defer recovery()

		var v float64
		switch value.(type) {
		case int, int8, int16, int32, int64:
			v = float64(reflect.ValueOf(value).Int())
		case uint, uint8, uint16, uint32, uint64:
			v = float64(reflect.ValueOf(value).Uint())
		case float32, float64:
			v = reflect.ValueOf(value).Float()
		default:
			return false
		}

		var a float64
		switch arg.(type) {
		case int, int8, int16, int32, int64:
			a = float64(reflect.ValueOf(arg).Int())
		case uint, uint8, uint16, uint32, uint64:
			a = float64(reflect.ValueOf(arg).Uint())
		case float32, float64:
			a = reflect.ValueOf(arg).Float()
		default:
			return false
		}

		return math.Mod(v, a) == 0
	},
	"lengthis": func(arg int, value interface{}) bool {
		defer recovery()

		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			return v.Len() == arg
		case reflect.String:
			return len([]rune(v.String())) == arg
		}

		return false
	},
	"trim": func(s string) string {
		defer recovery()

		return strings.TrimSpace(s)
	},
	"capfirst": func(s string) string {
		defer recovery()

		return strings.ToUpper(string(s[0])) + s[1:]
	},
	"pluralize": func(arg string, value interface{}) string {
		defer recovery()

		flag := false
		switch value.(type) {
		case int, int8, int16, int32, int64:
			flag = reflect.ValueOf(value).Int() == 1
		case uint, uint8, uint16, uint32, uint64:
			flag = reflect.ValueOf(value).Uint() == 1
		default:
			return ""
		}

		if !strings.Contains(arg, ",") {
			arg = "," + arg
		}

		bits := strings.Split(arg, ",")

		if len(bits) > 2 {
			return ""
		}

		if flag {
			return bits[0]
		}

		return bits[1]
	},
	"yesno": func(yes string, no string, value bool) string {
		defer recovery()

		if value {
			return yes
		}

		return no
	},
	"rjust": func(arg int, value string) string {
		defer recovery()

		n := arg - len([]rune(value))

		if n > 0 {
			value = strings.Repeat(" ", n) + value
		}

		return value
	},
	"ljust": func(arg int, value string) string {
		defer recovery()

		n := arg - len([]rune(value))

		if n > 0 {
			value = value + strings.Repeat(" ", n)
		}

		return value
	},
	"center": func(arg int, value string) string {
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
