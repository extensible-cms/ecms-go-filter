package string

import (
	"fmt"
	"testing"
)

type testCaseForString struct {
	Name     string
	Value    string
	Expected string
}

func ExpectEqual(tx *testing.T, a interface{}, b interface{}) {
	if a != b {
		tx.Errorf("Expected: %v;  Received: %v", b, a)
	}
}

func TestLowerCase(t *testing.T) {
	for _, tc := range []testCaseForString{
		{Value: "ABC", Expected: "abc"},
		{Value: "aEiOu", Expected: "aeiou"},
		{Value: "AeIoU", Expected: "aeiou"},
		{Value: "aeiou", Expected: "aeiou"},
	} {
		testName := fmt.Sprintf("LowerCase(%v) === %v", tc.Value, tc.Expected)
		t.Run(testName, func(t2 *testing.T) {
			result := LowerCase(tc.Value)
			ExpectEqual(t2, result, tc.Expected)
		})
	}
}

func TestTrim(t *testing.T) {
	for _, tc := range []testCaseForString{
		{Value: "    ", Expected: ""},
		{Value: "\t\t\t\t", Expected: ""},
		{Value: "\t\r\n\f", Expected: ""},
		{Value: " ABC ", Expected: "ABC"},
		{Value: "aEiOu", Expected: "aEiOu"},
		{Value: "   AeIoU", Expected: "AeIoU"},
		{Value: "AeIoU   ", Expected: "AeIoU"},
	} {
		testName := fmt.Sprintf("Trim(%v) === %v", tc.Value, tc.Expected)
		t.Run(testName, func(t2 *testing.T) {
			result := Trim(tc.Value)
			ExpectEqual(t2, result, tc.Expected)
		})
	}
}

func TestXmlEntities(t *testing.T) {
	type TestCase struct {
		Value    interface{}
		Expected string
	}
	testCases := make([]TestCase, 0)
	for k, v := range xmlEntitiesCharMap {
		strK := string(k)
		strV := string(v)
		testCases = append(testCases,
			TestCase{
				Value:    k,
				Expected: strV,
			},
			TestCase{
				Value:    byte(k),
				Expected: strV,
			},
			TestCase{
				Value:    strK + "abc",
				Expected: strV + "abc",
			},
			TestCase{
				Value:    []byte(strK + "abc"),
				Expected: strV + "abc",
			},
			TestCase{
				Value:    "abc" + strK,
				Expected: "abc" + strV,
			},
			TestCase{
				Value:    []byte("abc" + strK),
				Expected: "abc" + strV,
			},
			TestCase{
				Value:    "abc" + strK + "abc",
				Expected: "abc" + strV + "abc",
			},
			TestCase{
				Value:    []byte("abc" + strK + "abc"),
				Expected: "abc" + strV + "abc",
			},
			TestCase{
				Value:    "abc" + strK + "abc" + strK,
				Expected: "abc" + strV + "abc" + strV,
			},
			TestCase{
				Value:    []byte("abc" + strK + "abc" + strK),
				Expected: "abc" + strV + "abc" + strV,
			},
			TestCase{
				Value:    strK + "abc" + strK + "abc" + strK,
				Expected: strV + "abc" + strV + "abc" + strV,
			},
			TestCase{
				Value:    []byte(strK + "abc" + strK + "abc" + strK),
				Expected: strV + "abc" + strV + "abc" + strV,
			},
		)
	}

	testCases = append(testCases, TestCase{
		Value:    "<script>alert(\"bad script' here & should be escaped.\")</script>",
		Expected: "&lt;script&gt;alert(&quot;bad script&apos; here &amp; should be escaped.&quot;)&lt;/script&gt;",
	})

	for _, tc := range testCases {
		var given string
		switch tc.Value.(type) {
		case []byte:
			given = string(tc.Value.([]byte))
		case byte:
			given = string(tc.Value.(byte))
		default:
			given = tc.Value.(string)
		}
		testName := fmt.Sprintf("XmlEntities(%v) === %v", given, tc.Expected)
		t.Run(testName, func(t2 *testing.T) {
			result := XmlEntities(tc.Value)
			ExpectEqual(t2, result, tc.Expected)
		})
	}
}
