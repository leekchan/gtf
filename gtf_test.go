package gtf

import (
	"testing"
	"bytes"
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
	
	ParseTest(&buffer, "{{ \"안녕하세요. 반갑습니다.\" | stringTruncateChars 12 }}")
	AssertEqual(t, &buffer, "안녕하세요. 반갑...")
	
	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringTruncateChars 12 }}")
	AssertEqual(t, &buffer, "The Go Pr...")
	
	ParseTest(&buffer, "{{ \"안녕하세요. The Go Programming Language\" | stringTruncateChars 30 }}")
	AssertEqual(t, &buffer, "안녕하세요. The Go Programming L...")
	
	ParseTest(&buffer, "{{ \"The\" | stringTruncateChars 30 }}")
	AssertEqual(t, &buffer, "The")
	
	ParseTest(&buffer, "{{ \"The Go Programming Language\" | stringTruncateChars 3 }}")
	AssertEqual(t, &buffer, "The")
	
	ParseTest(&buffer, "{{ \"The Go\" | stringTruncateChars 6 }}")
	AssertEqual(t, &buffer, "The Go")
	
	ParseTest(&buffer, "{{ \"The Go\" | stringTruncateChars 30 }}")
	AssertEqual(t, &buffer, "The Go")
	
	ParseTest(&buffer, "{{ \"The Go\" | stringTruncateChars 0 }}")
	AssertEqual(t, &buffer, "")
	
	ParseTest(&buffer, "{{ \"The Go\" | stringTruncateChars -1 }}")
	AssertEqual(t, &buffer, "The Go")
}