package string

import (
	"bytes"
	"fmt"
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	bytes2 "github.com/extensible-cms/ecms-go-filter/bytes"
	"strings"
	"testing"
)

var nameSubSequences [][]byte

func init() {
	nameSubSequences = make([][]byte, 0)

	// Ensure only valid names are entered into names list
	for _, name := range bytes2.SubSequences([]byte("a-b_c:d")) { // subsequences for valid html names
		nameLastInd := len(name) - 1
		if len(name) == 0 ||
			bytes.Index(name, []byte{'-'}) == 0 || // skip on names starting with non-alpha-char
			bytes.Index(name, []byte{'_'}) == 0 || // ""
			bytes.Index(name, []byte{':'}) == 0 || // ""
			bytes.Index(name, []byte{'-'}) == nameLastInd || // skip on names ending with non-alpha-char
			bytes.Index(name, []byte{'_'}) == nameLastInd || //
			bytes.Index(name, []byte{':'}) == nameLastInd { //
			continue
		}
		nameSubSequences = append(nameSubSequences, name)
	}
}

func TestGetStripHtmlTags(t *testing.T) {
	type testCase struct {
		cases map[string]string
		names [][]byte
	}

	testCases := make([]testCase, 0)

	makeHtmlStr := func(s string) string {
		return fmt.Sprintf(
			"<!DOCTYPE html><html><head><title>Random Markup</title></head><body>%v</body></html>", s,
		)
	}

	randomHtmlContent := ""

	for _, tagName := range []string{"a", "b", "p", "i"} {
		randomHtmlContent = randomHtmlContent + fmt.Sprintf("<%v>random content</%v>", tagName, tagName)
		randomHtmlContent = randomHtmlContent + fmt.Sprintf(
			"<%v>random <%v>content</%></%v>", tagName, tagName, tagName, tagName,
		)
		randomHtmlContent = randomHtmlContent + fmt.Sprintf(
			"<%v>random <%v>content<%v><%v /></%v></%v></%v>",
			tagName, tagName, tagName, tagName, tagName, tagName, tagName,
		)
	}

	for _, name := range nameSubSequences {
		n := string(name)
		randomContent := "random content"
		attribWithRandom := n + "=\"" + randomContent + "\""
		attribWithSelf := n + "=\"" + n + "\""

		cases := map[string]string{
			"<div " + attribWithSelf + ">":                                            "<div " + attribWithSelf + ">",
			"<div " + attribWithSelf + " class=\"hello\">":                            "<div " + attribWithSelf + " class=\"hello\">",
			"<div " + attribWithRandom + " class=\"hello\" " + attribWithRandom + ">": "<div " + attribWithRandom + " class=\"hello\" " + attribWithRandom + ">",
			"<div class=\"hello\" " + attribWithRandom + " data-hello=\"hello\">":     "<div class=\"hello\" " + attribWithRandom + " data-hello=\"hello\">",
			"<div class=\"" + n + "\">":                                               "<div class=\"" + n + "\">",
			"":                                                                        "",
			// @todo add json value case
		}
		testCases = append(testCases, testCase{
			cases: cases,
			names: [][]byte{name},
		})
	}

	// @todo table-lize this suite
	// @todo add self closing tag case
	// @todo add case with self closing tag containing multiple attribs
	// @todo add case with opening tag containing multiple attribs
	// @todo add case with attributes containing json values
	for _, tagName := range nameSubSequences {
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
							t3.Errorf("Expected result not to contain removed tag %s;  Received %s", htmlTag, result)
						}
					}
				})

			}) // html cases loop
		}

	} // test cases loop
}

func TestGetStripPopulatedHtmlAttribs(t *testing.T) {
	type testCase struct {
		cases map[string]string
		names [][]byte
	}

	testCases := make([]testCase, 0)

	for _, name := range nameSubSequences {
		n := string(name)
		randomContent := "random content"
		attribWithRandom := n + "=\"" + randomContent + "\""
		attribWithSelf := n + "=\"" + n + "\""
		cases := map[string]string{
			// Self defining attributes (for now) will be ignored
			//"<div " + n + ">":                                           "<div>",
			//"<div " + n + " class=\"hello\">":                           "<div class=\"hello\">",
			"<div " + attribWithSelf + ">":                                            "<div>",
			"<div " + attribWithSelf + " class=\"hello\">":                            "<div class=\"hello\">",
			"<div " + attribWithRandom + " class=\"hello\" " + attribWithRandom + ">": "<div class=\"hello\">",
			"<div class=\"hello\" " + attribWithRandom + " data-hello=\"hello\">":     "<div class=\"hello\" data-hello=\"hello\">",
			"<div class=\"" + n + "\">":                                               "<div class=\"" + n + "\">",
			"":                                                                        "",
			// @todo add json value case
		}
		testCases = append(testCases, testCase{
			cases: cases,
			names: [][]byte{name},
		})
	}

	// Table-lize this suite
	for _, tc := range testCases {
		for attribCase, expectedValue := range tc.cases {
			testName := fmt.Sprintf("GetStripAttribTags(%s)(%s)",
				tc.names, attribCase,
			)

			t.Run(testName, func(t2 *testing.T) {
				f := GetStripPopulatedHtmlAttribs(tc.names)
				t2.Run("Expect `Filter` function", func(t3 *testing.T) {
					if f == nil {
						t3.Error("Expected a function;  Received `nil`")
					}
				})

				t2.Run(fmt.Sprintf("Expect attributes %s removed", tc.names), func(t3 *testing.T) {
					r := f(attribCase)
					if string(r.([]byte)) != expectedValue {
						t3.Errorf("Expected %s;  Received %s", expectedValue, r)
					}
				})

			})

		}

	} // test cases loop
}
