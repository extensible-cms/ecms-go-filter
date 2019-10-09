package string

import (
	"bytes"
	"fmt"
	bytes2 "github.com/extensible-cms/ecms-go-filter/bytes"
	"regexp"
	"testing"
)

var nameSubSequences [][]byte

// Generate a set of subsequences for given valid html tag names and valid html attributes (used by
// 	`TestStripHtmlTags` and `TestStripHtmlAttribs`).
func init() {
	// Capture subsequences here
	nameSubSequences = make([][]byte, 0)


 	rx := regexp.MustCompile("(:?(^[\\-:_])|([\\-:_]$))")

	// Ensure only valid names are entered into names list
	for _, name := range bytes2.SubSequences([]byte("a-b_c:d")) { // subsequences for valid html names

		// If empty name or invalid name start or end skip
		if len(name) == 0 || rx.Match(name) {
			continue
		}

		// Push name into captured names
		nameSubSequences = append(nameSubSequences, name)
	}
}

func TestGetStripHtmlTags(t *testing.T) {
	type testCase struct {
		cases map[string]string
		names [][]byte
	}

	testCases := make([]testCase, 0)

	// Random content to reuse in our test case permutations
	randomHtmlContent := ""
	randomContent := "Random content."

	// For random tag name create `random html markup and content`
	for _, tagName := range []string{"form", "input", "p", "i"} {
		randomHtmlContent = randomHtmlContent + fmt.Sprintf("<%v>random content</%v>", tagName, tagName)
		randomHtmlContent = randomHtmlContent + fmt.Sprintf(
			"<%v>random <%v>content</%v></%v>", tagName, tagName, tagName, tagName,
		)
		randomHtmlContent = randomHtmlContent + fmt.Sprintf(
			"<%v>random <%v>content{{deepInContent}}<%v><%v /></%v></%v></%v>",
			tagName, tagName, tagName, tagName, tagName, tagName, tagName,
		)
	}

	// For each in name subsequences create a test case
	for _, name := range nameSubSequences {
		n := string(name)
		openTag := fmt.Sprintf("<%s>", name)
		closeTag := fmt.Sprintf("</%s>", name)
		openTagWithAttribsAndSelf := fmt.Sprintf("<%s data-hello=\"world\" class=\"%s\">", name, name)
		openTagWithAttribs := fmt.Sprintf("<%s data-hello=\"world\" class=\"some-class-here\">", name)
		tagWithEmptyContent := openTag + closeTag
		tagWithContent := openTag + randomHtmlContent + closeTag
		tagWithAndSurroundingContent := randomHtmlContent + openTag + randomHtmlContent + closeTag + randomHtmlContent

		attribWithRandom := n + "=\"" + randomContent + "\""
		attribWithSelf := n + "=\"" + n + "\""
		//selfClosingTag := fmt.Sprintf("<%s />", name)
		//selfClosingTagWithAttribs := fmt.Sprintf("<%s %s %s %s %s />", "data-hi=\"hola\"",
		//	name, attribWithRandom, attribWithSelf, "class=\"hello-world\"",
		//)

		// Individual test case for current tag-name case
		cases := map[string]string{
			openTag: "",
			openTagWithAttribs: "",
			openTagWithAttribsAndSelf: "",
			closeTag: "",
			tagWithEmptyContent: "",
			tagWithContent: randomHtmlContent,
			tagWithAndSurroundingContent: randomHtmlContent + randomHtmlContent + randomHtmlContent,

			//selfClosingTagWithAttribs: "",
			//selfClosingTag: "",
			"<div " + attribWithSelf + ">":                                            "<div " + attribWithSelf + ">",
			"<div " + attribWithSelf + " class=\"hello\">":                            "<div " + attribWithSelf + " class=\"hello\">",
			"<div " + attribWithRandom + " class=\"hello\" " + attribWithRandom + ">": "<div " + attribWithRandom + " class=\"hello\" " + attribWithRandom + ">",
			"<div class=\"hello\" " + attribWithRandom + " data-hello=\"hello\">":     "<div class=\"hello\" " + attribWithRandom + " data-hello=\"hello\">",
			"<div class=\"" + n + "\">":                                               "<div class=\"" + n + "\">",
			"":                                                                        "",
			// @todo add json value case
		}

		// Append test case to test cases
		testCases = append(testCases, testCase{
			cases: cases,
			names: [][]byte{name},
		})
	}

	// @todo add self closing tag case
	// @todo add case with self closing tag containing multiple attribs
	// @todo add case with attributes containing json values
	// Walk through the test cases
	for _, tc := range testCases {
		for subject, expected := range tc.cases {

			// Get test cases 'set' name
			testName := fmt.Sprintf("GetStripHtmlTags(%s)(%s)", tc.names, subject)

			t.Run(testName, func(t2 *testing.T) {
				f := GetStripHtmlTags(tc.names)
				result := f(subject).([]byte)

				t2.Run("Expect `Filter` function", func(t3 *testing.T) {
					if f == nil {
						t3.Error("Expected a function;  Received `nil`")
					}
				})

				t2.Run(fmt.Sprintf("Expect %s", result), func(t3 *testing.T) {
					if false == bytes.Equal(result, []byte(expected)) {
						t3.Errorf("\nExpected \n%s;\nReceived: \n%s", expected, result)
						t3.Logf("Result length: %v\n", len(result))
					}
				})

			}) // cases loop

		} // sub test-cases map loop

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
