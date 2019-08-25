package string

import (
	"fmt"
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"math/rand"
	"strings"
	"testing"
)

func genRanChar() rune {
	return rand.Int31n(0x10FFFF)
}

func genRanRunes(max int32) []rune {
	var (
		s     []rune
		i     int32
		limit = rand.Int31n(max)
	)
	for i = 0; i < limit; i += 1 {
		s = append(s, genRanChar())
	}
	return s
}

func TestGetStripHtmlTagsFilter(t *testing.T) {
	for _, tagName := range ecmsGoFilter.StrSubSequences("abc_-:") { // subsequences for valid html symbol names
		if len(tagName) == 0 || strings.Index(tagName, "-") == 0 || // skip on tagNames starting with non-alpha-char
			strings.Index(tagName, "_") == 0 || strings.Index(tagName, ":") == 0 {
			continue
		}
		startTag := fmt.Sprintf("<%v>", tagName)
		closeTag := fmt.Sprintf("</%v>", tagName)
		randomContent := string(genRanRunes(100))
		htmlCases := ecmsGoFilter.StrSliceSubSequences([]string{
			startTag, closeTag, randomContent,
		})

		for _, htmlCase := range htmlCases {
			joinedHtml := strings.Join(htmlCase, "")
			testName := fmt.Sprintf("GetStripHtmlTags(%v)(%v)",
				tagName, joinedHtml,
			)
			t.Run(testName, func(t2 *testing.T) {
				f := GetStripHtmlTagsFilter([][]byte{[]byte(tagName)})
				result := f(joinedHtml)
				for _, htmlTag := range []string{startTag, closeTag} {
					if strings.Index(string(result.([]byte)), htmlTag) >= 0 {
						t2.Errorf("@todo add error message here")
					}
				}
			})
		}
	}
	//result := GetStripHtmlTagsFilter([][]byte{[]byte(tagName)})
}
