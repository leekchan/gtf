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

func ParseTest(buffer *bytes.Buffer, body string) {
	tpl := New("test").Funcs(GtfFuncMap)
	tpl.Parse(body)
	tpl.Execute(buffer, "")
}

func TestGtfFuncMap(t *testing.T) {
	var buffer bytes.Buffer

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringReplace \" \" }}")
	AssertEqual(t, &buffer, "TheGoProgrammingLanguage")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringDefault \"default value\" }}")
	AssertEqual(t, &buffer, "The Go Programming Language")

	ParseTest(&buffer, "{{ \"\" | stringDefault \"default value\" }}")
	AssertEqual(t, &buffer, "default value")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringLength }}")
	AssertEqual(t, &buffer, "27")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringLower }}")
	AssertEqual(t, &buffer, "the go programming language")
	
	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringUpper }}")
	AssertEqual(t, &buffer, "THE GO PROGRAMMING LANGUAGE")

	ParseTest(&buffer, "{{ \"안녕하세요. 반갑습니다.\" | stringTruncatechars 12 }}")
	AssertEqual(t, &buffer, "안녕하세요. 반갑...")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringTruncatechars 12 }}")
	AssertEqual(t, &buffer, "The Go Pr...")

	ParseTest(&buffer, "{{ \"안녕하세요. The Go Programming Language\" | stringTruncatechars 30 }}")
	AssertEqual(t, &buffer, "안녕하세요. The Go Programming L...")

	ParseTest(&buffer, "{{ \"The\" | stringTruncatechars 30 }}")
	AssertEqual(t, &buffer, "The")

	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringTruncatechars 3 }}")
	AssertEqual(t, &buffer, "The")

	ParseTest(&buffer, "{{ \"The Go\" | stringTruncatechars 6 }}")
	AssertEqual(t, &buffer, "The Go")

	ParseTest(&buffer, "{{ \"The Go\" | stringTruncatechars 30 }}")
	AssertEqual(t, &buffer, "The Go")

	ParseTest(&buffer, "{{ \"The Go\" | stringTruncatechars 0 }}")
	AssertEqual(t, &buffer, "")

	ParseTest(&buffer, "{{ \"The Go\" | stringTruncatechars -1 }}")
	AssertEqual(t, &buffer, "The Go")
	
	ParseTest(&buffer, "{{ \"http://www.example.org/foo?a=b&c=d\" | stringUrlencode }}")
	AssertEqual(t, &buffer, "http%3A%2F%2Fwww.example.org%2Ffoo%3Fa%3Db%26c%3Dd")
	
	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringWordcount }}")
	AssertEqual(t, &buffer, "4")
	
	ParseTest(&buffer, "{{ \"      The      Go       Programming      Language        \" | stringWordcount }}")
	AssertEqual(t, &buffer, "4")
	
	ParseTest(&buffer, "{{ 21 | intDivisibleby 3 }}")
	AssertEqual(t, &buffer, "true")
	
	ParseTest(&buffer, "{{ 21 | intDivisibleby 4 }}")
	AssertEqual(t, &buffer, "false")
	
	ParseTest(&buffer, "{{ \"Go\" | stringLengthIs 2 }}")
	AssertEqual(t, &buffer, "true")
	
	ParseTest(&buffer, "{{ \"안녕하세요.\" | stringLengthIs 6 }}")
	AssertEqual(t, &buffer, "true")
	
	ParseTest(&buffer, "{{ \"안녕하세요. Go!\" | stringLengthIs 10 }}")
	AssertEqual(t, &buffer, "true")
	
	ParseTest(&buffer, "{{ \"       The Go Programming Language     \" | stringTrim }}")
	AssertEqual(t, &buffer, "The Go Programming Language")
	
	ParseTest(&buffer, "{{ \"the go programming language\" | stringCapfirst }}")
	AssertEqual(t, &buffer, "The go programming language")
	
	ParseTest(&buffer, "You have 0 message{{ 0 | intPluralize \"s\" }}")
	AssertEqual(t, &buffer, "You have 0 messages")
	
	ParseTest(&buffer, "You have 1 message{{ 1 | intPluralize \"s\" }}")
	AssertEqual(t, &buffer, "You have 1 message")
	
	ParseTest(&buffer, "0 cand{{ 0 | intPluralize \"y,ies\" }}")
	AssertEqual(t, &buffer, "0 candies")
	
	ParseTest(&buffer, "1 cand{{ 1 | intPluralize \"y,ies\" }}")
	AssertEqual(t, &buffer, "1 candy")
	
	ParseTest(&buffer, "2 cand{{ 2 | intPluralize \"y,ies\" }}")
	AssertEqual(t, &buffer, "2 candies")
}
