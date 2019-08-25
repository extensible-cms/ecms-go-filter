package string

import (
	"bytes"
	"fmt"
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	bytes2 "github.com/extensible-cms/ecms-go-filter/bytes"
	"strings"
	"testing"
)

func TestGetStripHtmlTagsFilter(t *testing.T) {
	for _, tagName := range bytes2.SubSequences([]byte("a-b_c:d")) { // subsequences for valid html symbol names
		tagNameLastInd := len(tagName) - 1
		if len(tagName) == 0 ||
			bytes.Index(tagName, []byte{'-'}) == 0 || // skip on tagNames starting with non-alpha-char
			bytes.Index(tagName, []byte{'_'}) == 0 || // ""
			bytes.Index(tagName, []byte{':'}) == 0 || // ""
			bytes.Index(tagName, []byte{'-'}) == tagNameLastInd ||
			bytes.Index(tagName, []byte{'_'}) == tagNameLastInd ||
			bytes.Index(tagName, []byte{':'}) == tagNameLastInd {
			continue
		}

		openTag := fmt.Sprintf("<%s>", tagName)
		closeTag := fmt.Sprintf("</%s>", tagName)
		randomContent := "random content"

		// Random orderings of start, close and content nodes for our markup
		htmlCases := ecmsGoFilter.StrSliceSubSequences([]string{
			openTag, closeTag, randomContent,
		})

		for _, htmlCase := range htmlCases {
			joinedHtml := strings.Join(htmlCase, "")
			testName := fmt.Sprintf("GetStripHtmlTags(%s)(%s)",
				tagName, joinedHtml,
			)

			t.Run(testName, func(t2 *testing.T) {
				f := GetStripHtmlTags([][]byte{tagName})
				result := f(joinedHtml)

				t2.Run("Expect `Filter` function", func(t3 *testing.T) {
					if f == nil {
						t3.Error("Expected a function;  Received `nil`")
					}
				})

				t2.Run(fmt.Sprintf("Expect tags %s, %s removed", openTag, closeTag), func(t3 *testing.T) {
					for _, htmlTag := range []string{openTag, closeTag} {
						if bytes.Index(result.([]byte), []byte(htmlTag)) >= 0 {
							t2.Errorf("Expected result not to contain removed tag %s;  Received %s", htmlTag, result)
						}
					}
				})

			}) // html cases loop
		}

	} // test cases loop
}
