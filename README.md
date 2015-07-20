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
		filesize := 554832114
		tpl, _ := gtf.New("test").Parse("{{ . | filesizeformat }}")
		tpl.Execute(w, filesize)
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
		filesize := 554832114
		tpl, _ := template.New("test").Funcs(gtf.GtfFuncMap).Parse("{{ . | filesizeformat }}")
		tpl.Execute(w, filesize)
	})
    http.ListenAndServe(":8080", nil)
}
```



## Integration
You can use gtf with any web frameworks (revel, beego, martini, gin, etc) which use the Golang's built-in [html/template package](http://golang.org/pkg/html/template/). Calling ".Funcs(gtf.GtfFuncMap)" on [template.Template](http://golang.org/pkg/text/template/#Template) will add gtf functions to your template. I will add the detailed integration guides for each web framework soon!



## Safety
All gtf functions have their own recovery logics. The basic behavior of the recovery logic is silently swallowing all unexpected panics. All gtf functions would not make any panics in runtime. (**Production Ready!**)

If a panic occurs inside a gtf function, the function will silently swallow the panic and return "" (empty string). If you meet any unexpected empty output, [please make an issue](https://github.com/leekchan/gtf/issues/new)! :)



## Reference

### Index

* [replace](#replace)
* [default](#default)
* [length](#length)
* [lower](#lower)
* [upper](#upper)
* [truncatechars](#truncatechars)
* [urlencode](#urlencode)
* [wordcount](#wordcount)
* [divisibleby](#divisibleby)
* [lengthis](#lengthis)
* [trim](#trim)
* [capfirst](#capfirst)
* [pluralize](#pluralize)
* [yesno](#yesno)
* [rjust](#rjust)
* [ljust](#ljust)
* [center](#center)
* [filesizeformat](#filesizeformat)
* [apnumber](#apnumber)
* [intcomma](#intcomma)
* [ordinal](#ordinal)
* [first](#first)
* [last](#last)
* [join](#join)
* [slice](#slice)



#### replace

Removes all values of arg from the given string.

* supported value types : string
* supported argument types : string

```
{{ value | replace " " }}
```
If value is "The Go Programming Language", the output will be "TheGoProgrammingLanguage".



#### default

1. If the given string is ""(empty string), uses the given default argument.
1. If the given array/slice/map is empty, uses the given default argument.
1. If the given boolean value is false, uses the given default argument.

* supported value types : string, array, slice, map, boolean
* supported argument types : all

```
{{ value | default "default value" }}
```
If value is ""(the empty string), the output will be "default value".



#### length

Returns the length of the given string/array/slice/map.

* supported value types : string, array, slice, map

This function also supports unicode strings.

```
{{ value | length }}
```
If value is "The Go Programming Language", the output will be 27.



#### lower

Converts the given string into all lowercase.

* supported value types : string

```
{{ value | lower }}
```
If value is "The Go Programming Language", the output will be "the go programming language".



#### upper

Converts the given string into all uppercase.

* supported value types : string

```
{{ value | upper }}
```
If value is "The Go Programming Language", the output will be "THE GO PROGRAMMING LANGUAGE".



#### truncatechars

Truncates the given string if it is longer than the specified number of characters. Truncated strings will end with a translatable ellipsis sequence ("...")

* supported value types : string

**Argument:** Number of characters to truncate to

This function also supports unicode strings.

```
{{ value | truncatechars 12 }}
```

**Examples**

1. If input is {{ "The Go Programming Language" | truncatechars 12 }}, the output will be "The Go Pr...". (basic string)
1. If input is {{ "안녕하세요. 반갑습니다." | truncatechars 12 }}, the output will be "안녕하세요. 반갑...". (unicode)
1. If input is {{ "안녕하세요. The Go Programming Language" | truncatechars 30 }}, the output will be "안녕하세요. The Go Programming L...". (unicode)
1. If input is {{ "The" | truncatechars 30 }}, the output will be "The". (If the length of the given string is shorter than the argument, the output will be the original string.)
1. If input is {{ "The Go Programming Language" | truncatechars 3 }}, the output will be "The". (If the argument is less than or equal to 3, the output will not contain "...".)
1. If input is {{ "The Go" | truncatechars -1 }}, the output will be "The Go". (If the argument is less than 0, the argument will be ignored.)



#### urlencode

Escapes the given string for use in a URL.

* supported value types : string

```
{{ value | urlencode }}
```

If value is "http://www.example.org/foo?a=b&c=d", the output will be "http%3A%2F%2Fwww.example.org%2Ffoo%3Fa%3Db%26c%3Dd".



#### wordcount

Returns the number of words.

* supported value types : string

```
{{ value | wordcount }}
```

If value is "The Go Programming Language", the output will be 4.



#### divisibleby

Returns true if the value is divisible by the argument.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64
* supported argument types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

```
{{ value | divisibleby 3 }}
```

**Examples**

1. If input is {{ 21 | divisibleby 3 }}, the output will be true.
1. If input is {{ 21 | divisibleby 4 }}, the output will be false.
1. If input is {{ 3.0 | divisibleby 1.5 }}, the output will be true.



#### lengthis

Returns true if the value's length is the argument, or false otherwise.

* supported value types : string, array, slice, map
* supported argument types : int

```
{{ value | lengthis 3 }}
```

This function also supports unicode strings.

**Examples**

1. If input is {{ "Go" | lengthis 2 }}, the output will be true.
1. If input is {{ "안녕하세요. Go!" | lengthis 10 }}, the output will be true.



#### trim

Strips leading and trailing whitespace. 

* supported value types : string

```
{{ value | trim }}
```



#### capfirst

Capitalizes the first character of the given string.

* supported value types : string

```
{{ value | capfirst }}
```

If value is "the go programming language", the output will be "The go programming language".



#### pluralize

Returns a plural suffix if the value is not 1. You can specify both a singular and plural suffix, separated by a comma.

**Argument:** singular and plural suffix. 

1. "s" --> specify a singular suffix.
2. "y,ies" --> specify both a singular and plural suffix.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
* supported argument types : string

```
{{ value | pluralize "s" }}
{{ value | pluralize "y,ies" }}
```

**Examples**

1. You have 0 message{{ 0 | pluralize "s" }} --> You have 0 messages
2. You have 1 message{{ 1 | pluralize "s" }} --> You have 1 message
3. 0 cand{{ 0 | pluralize "y,ies" }} --> 0 candies
4. 1 cand{{ 1 | pluralize "y,ies" }} --> 1 candy
5. 2 cand{{ 2 | pluralize "y,ies" }} --> 2 candies



#### yesno

Returns argument strings according to the given boolean value.

* supported value types : boolean
* supported argument types : string

**Argument:** any value for true and false

```
{{ value | yesno "yes!" "no!" }}
```


#### rjust

Right-aligns the given string in a field of a given width. This function also supports unicode strings. 

* supported value types : string

```
{{ value | rjust 10 }}
```

**Examples**

1. If input is {{ "Go" | rjust 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Go".
1. If input is {{ "안녕하세요" | rjust 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;안녕하세요".



#### ljust

Left-aligns the given string in a field of a given width. This function also supports unicode strings. 

* supported value types : string

```
{{ value | ljust 10 }}
```

**Examples**

1. If input is {{ "Go" | ljust 10 }}, the output will be "Go&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;".
1. If input is {{ "안녕하세요" | ljust 10 }}, the output will be "안녕하세요&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;".



#### center

Centers the given string in a field of a given width. This function also supports unicode strings. 

* supported value types : string

```
{{ value | center 10 }}
```

**Examples**

1. If input is {{ "Go" | center 10 }}, the output will be "&nbsp;&nbsp;&nbsp;&nbsp;Go&nbsp;&nbsp;&nbsp;&nbsp;".
1. If input is {{ "안녕하세요" | center 10 }}, the output will be "&nbsp;&nbsp;안녕하세요&nbsp;&nbsp;&nbsp;".



#### filesizeformat

Formats the value like a human readable file size.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64

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



#### apnumber

For numbers 1-9, returns the number spelled out. Otherwise, returns the number. 

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64

```
{{ value | apnumber }}
```

**Examples**

1. {{ 1 | apnumber }} --> one
1. {{ 2 | apnumber }} --> two
1. {{ 3 | apnumber }} --> three
1. {{ 9 | apnumber }} --> nine
1. {{ 10 | apnumber }} --> 10
1. {{ 1000 | apnumber }} --> 1000



#### intcomma

Converts an integer to a string containing commas every three digits.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64

```
{{ value | intcomma }}
```

**Examples**

1. {{ 1000 | intcomma }} --> 1,000
1. {{ -1000 | intcomma }} --> -1,000
1. {{ 1578652313 | intcomma }} --> 1,578,652,313



#### ordinal

Converts an integer to its ordinal as a string.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64

```
{{ value | ordinal }}
```

**Examples**

1. {{ 1 | ordinal }} --> 1st
1. {{ 2 | ordinal }} --> 2nd
1. {{ 3 | ordinal }} --> 3rd
1. {{ 11 | ordinal }} --> 11th
1. {{ 12 | ordinal }} --> 12th
1. {{ 13 | ordinal }} --> 13th
1. {{ 14 | ordinal }} --> 14th



#### first

Returns the first item in the given value.

* supported value types : string, slice, array

This function also supports unicode strings.

```
{{ value | first }}
```

**Examples**

1. If value is the string "The go programming language", the output will be the string "T".
1. If value is the string "안녕하세요", the output will be the string "안". (unicode)
1. If value is the slice []string{"go", "python", "ruby"}, the output will be the string "go".
1. If value is the array [3]string{"go", "python", "ruby"}, the output will be the string "go".



#### last

Returns the last item in the given value.

* supported value types : string, slice, array

This function also supports unicode strings.

```
{{ value | last }}
```

**Examples**

1. If value is the string "The go programming language", the output will be the string "e".
1. If value is the string "안녕하세요", the output will be the string "요". (unicode)
1. If value is the slice []string{"go", "python", "ruby"}, the output will be the string "ruby".
1. If value is the array [3]string{"go", "python", "ruby"}, the output will be the string "ruby".




#### join

Concatenates the given slice to create a single string. The given argument (separator) will be placed between elements in the resulting string.

```
{{ value | join " " }}
```

If value is the slice []string{"go", "python", "ruby"}, the output will be the string "go python ruby"




#### slice

Returns a slice of the given value. The first argument is the start position, and the second argument is the end position.

* supported value types : string, slice
* supported argument types : int

This function also supports unicode strings.

```
{{ value | slice 0 2 }}
```

**Examples**

1. If input is {{ "The go programming language" | slice 0 6 }}, the output will be "The go".
1. If input is {{ "안녕하세요" | slice 0 2 }}, the output will be "안녕". (unicode)
1. If input is {{ []string{"go", "python", "ruby"} | slice 0 2 }}, the output will be []string{"go", "python"}.



## Goal
The first goal is implementing all built-in template filters of Django & Jinja2.

* [Django | Built-in filter reference](https://docs.djangoproject.com/en/1.8/ref/templates/builtins/#built-in-filter-reference)
* [Jinja2 | List of Builtin Filters](http://jinja.pocoo.org/docs/dev/templates/#builtin-filters)

The final goal is building a ultimate set which contains hundreds of useful template functions.


## Contributing
I love pull requests :) You can add any useful template functions by submitting a pull request!
