package data

import (
	"bytes"
	"strings"
	"testing"
)

func TestStripTags(t *testing.T) {

	var html = `<html>
	<head>
	    <title>Test</title>
	</head>
	<body>
	<p>Hello. My name is james</p>
	<p>This is a test to see if the thing will work</p>
	</body>
	</html>`

	//t.Log("Stripping")
	docText := StripHTML(bytes.NewBufferString(html))

	if !strings.Contains(docText, "Hello. My name is james") {
		t.Error("Document text returned is not complete")
		t.Fail()
	}

	if strings.ContainsAny(docText, "<>") {
		t.Error("Document contains HTML - epic fail")
		t.Fail()
	}

}
