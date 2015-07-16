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

### Method 2 : Adds gtf functions to the existing template.

You can also add the gtf functions to the existing template. Just call ".Funcs(gtf.GtfFuncMap)".

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


## Safety
All gtf functions have their own recovery logics. The basic behavior of the recovery logic is silently swallowing all unexpected panics. All gtf functions would not make any panics in runtime. (**Production Ready!**)

If a panic occurs inside a gtf function, the function will silently swallow the panic and return "" (empty string). If you meet any unexpected empty output, [please make an issue](https://github.com/leekchan/gtf/issues/new)! :)



## Naming convention

1. prefix(type of input value) + function : It supports only one type.
2. no prefix + function : It supports various types and evaluates input type in runtime. Please refer to "supported value types" of each function in reference. (For example, the input value of "filesizeformat" could be int or float.)


**Examples**

1. <a href="#stringLength">stringLength</a> => The type of the input value is string, and the function will return the length of the given value.
1. <a href="#intDivisibleby">intDivisibleby</a> => The type of the input value is int, and the function will return true if the value is divisible by the argument.
1. <a href="#filesizeformat">filesizeformat</a> => It supports various types(int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64).



## Reference

### Index

* <a href="#stringReplace">stringReplace</a>
* <a href="#stringDefault">stringDefault</a>
* <a href="#stringLength">stringLength</a>
* <a href="#stringLower">stringLower</a>
* <a href="#stringUpper">stringUpper</a>
* <a href="#stringTruncatechars">stringTruncatechars</a>
* <a href="#stringUrlencode">stringUrlencode</a>
* <a href="#stringWordcount">stringWordcount</a>
* <a href="#intDivisibleby">intDivisibleby</a>
* <a href="#stringLengthIs">stringLengthIs</a>
* <a href="#stringTrim">stringTrim</a>
* <a href="#stringCapfirst">stringCapfirst</a>
* <a href="#intPluralize">intPluralize</a>
* <a href="#boolYesno">boolYesno</a>
* <a href="#stringRjust">stringRjust</a>
* <a href="#stringLjust">stringLjust</a>
* <a href="#stringCenter">stringCenter</a>
* <a href="#filesizeformat">filesizeformat</a>

#### stringReplace

Removes all values of arg from the given string.

```
{{ value | stringReplace " " }}
```
If value is "The Go Programming Language", the output will be "TheGoProgrammingLanguage".



#### stringDefault

If value is ""(the empty string), uses the given default string.

```
{{ value | stringDefault "default value" }}
```
If value is ""(the empty string), the output will be "default value".



#### stringLength

Returns the length of the given string.

```
{{ value | stringLength }}
```
If value is "The Go Programming Language", the output will be 27.



#### stringLower

Converts the given string into all lowercase.

```
{{ value | stringLower }}
```
If value is "The Go Programming Language", the output will be "the go programming language".



#### stringUpper

Converts the given string into all uppercase.

```
{{ value | stringUpper }}
```
If value is "The Go Programming Language", the output will be "THE GO PROGRAMMING LANGUAGE".



#### stringTruncatechars

Truncates the given string if it is longer than the specified number of characters. Truncated strings will end with a translatable ellipsis sequence ("...")

**Argument:** Number of characters to truncate to

This function also supports unicode strings.

```
{{ value | stringTruncatechars 12 }}
```

**Examples**

1. If input is {{ "The Go Programming Language" | stringTruncatechars 12 }}, the output will be "The Go Pr...". (basic string)
1. If input is {{ "안녕하세요. 반갑습니다." | stringTruncatechars 12 }}, the output will be "안녕하세요. 반갑...". (unicode)
1. If input is {{ "안녕하세요. The Go Programming Language" | stringTruncatechars 30 }}, the output will be "안녕하세요. The Go Programming L...". (unicode)
1. If input is {{ "The" | stringTruncatechars 30 }}, the output will be "The". (If the length of the given string is shorter than the argument, the output will be the original string.)
1. If input is {{ "The Go Programming Language" | stringTruncatechars 3 }}, the output will be "The". (If the argument is less than or equal to 3, the output will not contain "...".)
1. If input is {{ "The Go" | stringTruncatechars -1 }}, the output will be "The Go". (If the argument is less than 0, the argument will be ignored.)



#### stringUrlencode

Escapes the given string for use in a URL.

```
{{ value | stringUrlencode }}
```

If value is "http://www.example.org/foo?a=b&c=d", the output will be "http%3A%2F%2Fwww.example.org%2Ffoo%3Fa%3Db%26c%3Dd".



#### stringWordcount

Returns the number of words.

```
{{ value | stringWordcount }}
```

If value is "The Go Programming Language", the output will be 4.



#### intDivisibleby

Returns true if the value is divisible by the argument.

```
{{ value | intDivisibleby 3 }}
```

**Examples**

1. If input is {{ 21 | intDivisibleby 3 }}, the output will be true.
1. If input is {{ 21 | intDivisibleby 4 }}, the output will be false.



#### stringLengthIs

Returns true if the value's length is the argument, or false otherwise.

```
{{ value | stringLengthIs 3 }}
```

This function also supports unicode strings.

**Examples**

1. If input is {{ "Go" | stringLengthIs 2 }}, the output will be true.
1. If input is {{ "안녕하세요. Go!" | stringLengthIs 10 }}, the output will be true.



#### stringTrim

Strips leading and trailing whitespace. 

```
{{ value | stringTrim }}
```



#### stringCapfirst

Capitalizes the first character of the given string.

```
{{ value | stringCapfirst }}
```

If value is "the go programming language", the output will be "The go programming language".



#### intPluralize

Returns a plural suffix if the value is not 1. You can specify both a singular and plural suffix, separated by a comma.

**Argument:** singular and plural suffix. 

1. "s" --> specify a singular suffix.
2. "y,ies" --> specify both a singular and plural suffix.

```
{{ value | intPluralize "s" }}
{{ value | intPluralize "y,ies" }}
```

**Examples**

1. You have 0 message{{ 0 | intPluralize "s" }} --> You have 0 messages
2. You have 1 message{{ 1 | intPluralize "s" }} --> You have 1 message
3. 0 cand{{ 0 | intPluralize "y,ies" }} --> 0 candies
4. 1 cand{{ 1 | intPluralize "y,ies" }} --> 1 candy
5. 2 cand{{ 2 | intPluralize "y,ies" }} --> 2 candies



#### boolYesno

Returns argument strings according to the given boolean value.

**Argument:** strings for true and false

```
{{ value | boolYesno "yes!" "no!" }}
```


#### stringRjust

Right-aligns the given string in a field of a given width. This function also supports unicode strings. 

```
{{ value | stringRjust 10 }}
```

**Examples**

1. If input is {{ "Go" | stringRjust 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Go".
1. If input is {{ "안녕하세요" | stringRjust 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;안녕하세요".



#### stringLjust

Left-aligns the given string in a field of a given width. This function also supports unicode strings. 

```
{{ value | stringLjust 10 }}
```

**Examples**

1. If input is {{ "Go" | stringLjust 10 }}, the output will be "Go&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;".
1. If input is {{ "안녕하세요" | stringLjust 10 }}, the output will be "안녕하세요&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;".



#### stringCenter

Centers the given string in a field of a given width. This function also supports unicode strings. 

```
{{ value | stringCenter 10 }}
```

**Examples**

1. If input is {{ "Go" | stringCenter 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;Go&nbsp;&nbsp;&nbsp;&nbsp;".
1. If input is {{ "안녕하세요" | stringCenter 10 }}, the output will be "&nbsp;&nbsp;안녕하세요&nbsp;&nbsp;&nbsp;".



#### filesizeformat

Formats the value like a human readable file size.

**supported value types:** int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```
{{ value | filesizeformat }}
```

**Examples**

1. {{ 234 | filesizeformat }} --> "234 bytes"
1. {{ 12345 | filesizeformat }} --> "12.1 KB"
1. {{ 12345.35335 | filesizeformat }} --> "12.1 KB"
1. {{ 1048576 | filesizeformat } --> "1 MB"
1. {{ 554832114 | filesizeformat }} --> "529.1 MB"
1. {{ 14868735121 | filesizeformat }} --> "13.8 GB"
1. {{ 14868735121365 | filesizeformat }} --> "13.5 TB"
1. {{ 1486873512136523 | filesizeformat }} --> "1.3 PB"



## Goal
The first goal is implementing all built-in template filters of Django & Jinja2.

* [Django | Built-in filter reference](https://docs.djangoproject.com/en/1.8/ref/templates/builtins/#built-in-filter-reference)
* [Jinja2 | List of Builtin Filters](http://jinja.pocoo.org/docs/dev/templates/#builtin-filters)

The final goal is building a ultimate set which contains hundreds of useful template functions.


## Contributing
I love pull requests :) You can add any useful template functions by submitting a pull request!
