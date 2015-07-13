package gtf

import (
	"testing"
	"bytes"
	"html/template"
)

func AssertEqual(t *testing.T, buffer *bytes.Buffer, testString string) {
	if buffer.String() != testString {
		t.Error()
	}
	buffer.Reset()
}

func ParseTest(buffer *bytes.Buffer, body string) {
	tpl := template.New("test").Funcs(GtfFuncMap)
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
	
	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringLower }}")
	AssertEqual(t, &buffer, "the go programming language")
}