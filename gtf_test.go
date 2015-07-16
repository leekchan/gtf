package gtf

import (
	"bytes"
	"testing"
)

func AssertEqual(t *testing.T, buffer *bytes.Buffer, testString string) {
	if buffer.String() != testString {
		t.Errorf("Expected %s, got %s", testString, buffer.String())
	}
	buffer.Reset()
}

func ParseTest(buffer *bytes.Buffer, body string, data interface{}) {
	tpl := New("test").Funcs(GtfFuncMap)
	tpl.Parse(body)
	tpl.Execute(buffer, data)
}

func TestGtfFuncMap(t *testing.T) {
	var buffer bytes.Buffer

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | replace \" \" }}", "")
	AssertEqual(t, &buffer, "TheGoProgrammingLanguage")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | default \"default value\" }}", "")
	AssertEqual(t, &buffer, "The Go Programming Language")

	ParseTest(&buffer, "{{ \"\" | default \"default value\" }}", "")
	AssertEqual(t, &buffer, "default value")

	ParseTest(&buffer, "{{ . | default \"default value\" }}", []string{"go", "python", "ruby"})
	AssertEqual(t, &buffer, "[go python ruby]")

	ParseTest(&buffer, "{{ . | default 10 }}", []int{})
	AssertEqual(t, &buffer, "10")
	
	ParseTest(&buffer, "{{ . | default \"empty\" }}", false)
	AssertEqual(t, &buffer, "empty")
	
	ParseTest(&buffer, "{{ . | default \"empty\" }}", 1)
	AssertEqual(t, &buffer, "1")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | length }}", "")
	AssertEqual(t, &buffer, "27")

	ParseTest(&buffer, "{{ \"안녕하세요\" | length }}", "")
	AssertEqual(t, &buffer, "5")

	ParseTest(&buffer, "{{ . | length }}", []string{"go", "python", "ruby"})
	AssertEqual(t, &buffer, "3")

	ParseTest(&buffer, "{{ . | length }}", false)
	AssertEqual(t, &buffer, "0")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | lower }}", "")
	AssertEqual(t, &buffer, "the go programming language")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | upper }}", "")
	AssertEqual(t, &buffer, "THE GO PROGRAMMING LANGUAGE")

	ParseTest(&buffer, "{{ \"안녕하세요. 반갑습니다.\" | truncatechars 12 }}", "")
	AssertEqual(t, &buffer, "안녕하세요. 반갑...")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | truncatechars 12 }}", "")
	AssertEqual(t, &buffer, "The Go Pr...")

	ParseTest(&buffer, "{{ \"안녕하세요. The Go Programming Language\" | truncatechars 30 }}", "")
	AssertEqual(t, &buffer, "안녕하세요. The Go Programming L...")

	ParseTest(&buffer, "{{ \"The\" | truncatechars 30 }}", "")
	AssertEqual(t, &buffer, "The")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | truncatechars 3 }}", "")
	AssertEqual(t, &buffer, "The")

	ParseTest(&buffer, "{{ \"The Go\" | truncatechars 6 }}", "")
	AssertEqual(t, &buffer, "The Go")

	ParseTest(&buffer, "{{ \"The Go\" | truncatechars 30 }}", "")
	AssertEqual(t, &buffer, "The Go")

	ParseTest(&buffer, "{{ \"The Go\" | truncatechars 0 }}", "")
	AssertEqual(t, &buffer, "")

	ParseTest(&buffer, "{{ \"The Go\" | truncatechars -1 }}", "")
	AssertEqual(t, &buffer, "The Go")

	ParseTest(&buffer, "{{ \"http://www.example.org/foo?a=b&c=d\" | urlencode }}", "")
	AssertEqual(t, &buffer, "http%3A%2F%2Fwww.example.org%2Ffoo%3Fa%3Db%26c%3Dd")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | wordcount }}", "")
	AssertEqual(t, &buffer, "4")

	ParseTest(&buffer, "{{ \"      The      Go       Programming      Language        \" | wordcount }}", "")
	AssertEqual(t, &buffer, "4")

	ParseTest(&buffer, "{{ 21 | divisibleby 3 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ 21 | divisibleby 4 }}", "")
	AssertEqual(t, &buffer, "false")

	ParseTest(&buffer, "{{ 3.0 | divisibleby 3 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ 3.0 | divisibleby 1.5 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ . | divisibleby 1.5 }}", uint(300))
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ 12 | divisibleby . }}", uint(3))
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ 21 | divisibleby 4 }}", "")
	AssertEqual(t, &buffer, "false")

	ParseTest(&buffer, "{{ false | divisibleby 3 }}", "")
	AssertEqual(t, &buffer, "false")

	ParseTest(&buffer, "{{ 3 | divisibleby false }}", "")
	AssertEqual(t, &buffer, "false")

	ParseTest(&buffer, "{{ \"Go\" | lengthis 2 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ \"안녕하세요.\" | lengthis 6 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ \"안녕하세요. Go!\" | lengthis 10 }}", "")
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ . | lengthis 3 }}", []string{"go", "python", "ruby"})
	AssertEqual(t, &buffer, "true")

	ParseTest(&buffer, "{{ . | lengthis 3 }}", false)
	AssertEqual(t, &buffer, "false")

	ParseTest(&buffer, "{{ \"       The Go Programming Language     \" | trim }}", "")
	AssertEqual(t, &buffer, "The Go Programming Language")

	ParseTest(&buffer, "{{ \"the go programming language\" | capfirst }}", "")
	AssertEqual(t, &buffer, "The go programming language")

	ParseTest(&buffer, "You have 0 message{{ 0 | pluralize \"s\" }}", "")
	AssertEqual(t, &buffer, "You have 0 messages")

	ParseTest(&buffer, "You have 1 message{{ 1 | pluralize \"s\" }}", "")
	AssertEqual(t, &buffer, "You have 1 message")

	ParseTest(&buffer, "0 cand{{ 0 | pluralize \"y,ies\" }}", "")
	AssertEqual(t, &buffer, "0 candies")

	ParseTest(&buffer, "1 cand{{ 1 | pluralize \"y,ies\" }}", "")
	AssertEqual(t, &buffer, "1 candy")

	ParseTest(&buffer, "2 cand{{ 2 | pluralize \"y,ies\" }}", "")
	AssertEqual(t, &buffer, "2 candies")

	ParseTest(&buffer, "{{ 2 | pluralize \"y,ies,s\" }}", "")
	AssertEqual(t, &buffer, "")

	ParseTest(&buffer, "2 cand{{ . | pluralize \"y,ies\" }}", uint(2))
	AssertEqual(t, &buffer, "2 candies")

	ParseTest(&buffer, "1 cand{{ . | pluralize \"y,ies\" }}", uint(1))
	AssertEqual(t, &buffer, "1 candy")

	ParseTest(&buffer, "{{ . | pluralize \"y,ies\" }}", "test")
	AssertEqual(t, &buffer, "")

	ParseTest(&buffer, "{{ true | yesno \"yes~\" \"no~\" }}", "")
	AssertEqual(t, &buffer, "yes~")

	ParseTest(&buffer, "{{ false | yesno \"yes~\" \"no~\" }}", "")
	AssertEqual(t, &buffer, "no~")

	ParseTest(&buffer, "{{ \"Go\" | rjust 10 }}", "")
	AssertEqual(t, &buffer, "        Go")

	ParseTest(&buffer, "{{ \"안녕하세요\" | rjust 10 }}", "")
	AssertEqual(t, &buffer, "     안녕하세요")

	ParseTest(&buffer, "{{ \"Go\" | ljust 10 }}", "")
	AssertEqual(t, &buffer, "Go        ")

	ParseTest(&buffer, "{{ \"안녕하세요\" | ljust 10 }}", "")
	AssertEqual(t, &buffer, "안녕하세요     ")

	ParseTest(&buffer, "{{ \"Go\" | center 10 }}", "")
	AssertEqual(t, &buffer, "    Go    ")

	ParseTest(&buffer, "{{ \"안녕하세요\" | center 10 }}", "")
	AssertEqual(t, &buffer, "  안녕하세요   ")

	ParseTest(&buffer, "{{ 123456789 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "117.7 MB")

	ParseTest(&buffer, "{{ 234 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "234 bytes")

	ParseTest(&buffer, "{{ 12345 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "12.1 KB")

	ParseTest(&buffer, "{{ 554832114 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "529.1 MB")

	ParseTest(&buffer, "{{ 1048576 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "1 MB")

	ParseTest(&buffer, "{{ 14868735121 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "13.8 GB")

	ParseTest(&buffer, "{{ 14868735121365 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "13.5 TB")

	ParseTest(&buffer, "{{ 1486873512136523 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "1.3 PB")

	ParseTest(&buffer, "{{ 12345.35335 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "12.1 KB")

	ParseTest(&buffer, "{{ 4294967293 | filesizeformat }}", "")
	AssertEqual(t, &buffer, "4 GB")

	ParseTest(&buffer, "{{ \"Go\" | filesizeformat }}", "")
	AssertEqual(t, &buffer, "")

	ParseTest(&buffer, "{{ . | filesizeformat }}", uint(500))
	AssertEqual(t, &buffer, "500 bytes")
}
