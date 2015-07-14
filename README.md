# gtf - a useful set of Golang Template Functions
[![Build Status](https://travis-ci.org/leekchan/gtf.svg?branch=master)](https://travis-ci.org/leekchan/gtf)
[![Coverage Status](https://coveralls.io/repos/leekchan/gtf/badge.svg?branch=master&service=github)](https://coveralls.io/github/leekchan/gtf?branch=master)
[![GoDoc](https://godoc.org/github.com/leekchan/gtf?status.svg)](https://godoc.org/github.com/leekchan/gtf)

gtf is a useful set of Golang Template Functions. The goal of this project is implementing all built-in template filters of Django & Jinja2. 

## Basic usages

### Method 1 : Uses gtf.New

gtf.New is a wrapper function of [template.New](http://golang.org/pkg/text/template/#New). It automatically adds the gtf functions to the template's function map and returns [template.Template](http://golang.org/pkg/text/template/#Template).

```Go
package main

import (
	"net/http"
	"github.com/leekchan/gtf"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := gtf.New("test").Parse("{{ \"The Go Programming Language\" | stringReplace \" \" }}")
		tpl.Execute(w, "")
	})
    http.ListenAndServe(":8080", nil)
}
```

### Method 2 : Adds gtf functions to the exsisting template.

You can also add the gtf functions to the exsisting template. Just call ".Funcs(gtf.GtfFuncMap)".

```Go
package main

import (
	"net/http"
	"html/template"
	"github.com/leekchan/gtf"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.New("test").Funcs(gtf.GtfFuncMap).Parse("{{ \"The Go Programming Language\" | stringReplace \" \" }}")
		tpl.Execute(w, "")
	})
    http.ListenAndServe(":8080", nil)
}
```

## Reference
### stringReplace

Removes all values of arg from the given string.

```
{{ value | stringReplace " " }}
```
If value is "The Go Programming Language", the output will be "TheGoProgrammingLanguage".



### stringDefault

If value is ""(the empty string), uses the given default string.

```
{{ value | stringDefault "default value" }}
```
If value is ""(the empty string), the output will be "default value".



### stringLength

Returns the length of the given string.

```
{{ value | stringLength }}
```
If value is "The Go Programming Language", the output will be 27.



### stringLower

Converts the given string into all lowercase.

```
{{ value | stringLower }}
```
If value is "The Go Programming Language", the output will be "the go programming language".



### stringTruncateChars

Truncates the given string if it is longer than the specified number of characters. Truncated strings will end with a translatable ellipsis sequence ("...")

**Argument:** Number of characters to truncate to

This function also supports unicode strings.

```
{{ value | stringTruncateChars:12 }}
```

**Examples**

1. If input is {{ "The Go Programming Language" | stringTruncateChars 12 }}, the output will be "The Go Pr...". (basic string)
1. If input is {{ "안녕하세요. 반갑습니다." | stringTruncateChars 12 }}, the output will be "안녕하세요. 반갑...". (unicode)
1. If input is {{ "안녕하세요. The Go Programming Language" | stringTruncateChars 30 }}, the output will be "안녕하세요. The Go Programming L...". (unicode)
1. If input is {{ "The" | stringTruncateChars 30 }}, the output will be "The". (If the length of the given string is shorter than the argument, the output will be the original string.)
1. If input is {{ "The Go Programming Language" | stringTruncateChars 3 }}, the output will be "The". (If the argument is less than or equal to 3, the output will not contain "...".)
1. If input is {{ "The Go" | stringTruncateChars -1 }}, the output will be "The Go". (If the argument is less than 0, the argument will be ignored.)



## Goal
The main goal is implementing all built-in template filters of Django & Jinja2.

* [Django | Built-in filter reference](https://docs.djangoproject.com/en/1.8/ref/templates/builtins/#built-in-filter-reference)
* [Jinja2 | List of Builtin Filters](http://jinja.pocoo.org/docs/dev/templates/#builtin-filters)