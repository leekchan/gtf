# gtf - a useful set of Golang Template Functions
[![Build Status](https://travis-ci.org/leekchan/gtf.svg?branch=master)](https://travis-ci.org/leekchan/gtf)
[![Coverage Status](https://coveralls.io/repos/leekchan/gtf/badge.svg?branch=master&service=github)](https://coveralls.io/github/leekchan/gtf?branch=master)

gtf is a useful set of Golang Template Functions. The goal of this project is implementing all built-in template filters of Django & Jinja2. 

## Basic Example

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

## Reference (TODO)
### stringReplace

Removes all values of arg from the given string.

```
{{ value | stringReplace " " }}
```
If value is "The Go Programming Language", the output will be "TheGoProgrammingLanguage"



### stringDefault

If value is ""(the empty string), uses the given default string.

```
{{ value | stringDefault "default value" }}
```
If value is ""(the empty string), the output will be "default value"



### stringLength

Returns the length of the given string.

```
{{ value | stringLength }}
```
If value is "The Go Programming Language", the output will be 27



### stringLower

Converts the given string into all lowercase.

```
{{ value | stringLower }}
```
If value is "The Go Programming Language", the output will be "the go programming language"



## Goal
The main goal is implementing all built-in template filters of Django & Jinja2.

* [Django | Built-in filter reference](https://docs.djangoproject.com/en/1.8/ref/templates/builtins/#built-in-filter-reference)
* [Jinja2 | List of Builtin Filters](http://jinja.pocoo.org/docs/dev/templates/#builtin-filters)