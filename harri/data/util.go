package data

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

/*StripHTML removes html tags from documents and just leaves text
 */
func StripHTML(htmlReader io.Reader) string {
	z := html.NewTokenizer(htmlReader)

	inAside := false
	inFigure := false

	var docText bytes.Buffer

	for {
		//token type
		tokenType := z.Next()
		token := z.Token()

		switch {

		// handle start/open tag
		case tokenType == html.StartTagToken:
			inAside = token.Data == "aside"
			inFigure = token.Data == "figure"
			break

		// handle end tag
		case tokenType == html.EndTagToken:

			if token.Data == "aside" {
				inAside = false
			}

			if token.Data == "figure" {
				inFigure = false
			}

			if token.Data == "p" {
				// append newlines to delineate new paragraph
				docText.WriteString("\n\n")
			}
			break

		// handle text/cdata content
		case tokenType == html.TextToken:

			if !(inAside || inFigure) {
				docText.WriteString(strings.TrimSpace(token.Data) + " ")
			}
			break

		case tokenType == html.ErrorToken:
			fmt.Printf(token.Data)
			// End of the document, we're done
			return fmt.Sprint(docText.String())
		}

	}

}
