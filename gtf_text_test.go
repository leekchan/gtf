package gtf

import (
	"bytes"
	"testing"
	"text/template"
)

func TextTemplateParseTest(buffer *bytes.Buffer, body string, data interface{}) {
	tpl := template.New("test").Funcs(GtfTextFuncMap)
	tpl.Parse(body)
	tpl.Execute(buffer, data)
}

func TestTextTemplateGtfFuncMap(t *testing.T) {
	var buffer bytes.Buffer

	TextTemplateParseTest(&buffer, "{{ \"The Go Programming Language\" | replace \" \" }}", "")
	AssertEqual(t, &buffer, "TheGoProgrammingLanguage")

	TextTemplateParseTest(&buffer, "{{ \"The Go Programming Language\" | length }}", "")
	AssertEqual(t, &buffer, "27")

	TextTemplateParseTest(&buffer, "{{ 21 | divisibleby 3 }}", "")
	AssertEqual(t, &buffer, "true")
}
