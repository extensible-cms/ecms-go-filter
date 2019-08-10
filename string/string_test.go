package string

import (
	"fmt"
	"testing"
)

type TestCaseForString struct {
	Name     string
	Value    string
	Expected string
}

func ExpectEqual(tx *testing.T, a interface{}, b interface{}) {
	if a != b {
		tx.Errorf("Expected: %v;  Received: %v", b, a)
	}
}

func TestToLowerCase(t *testing.T) {
	for _, tc := range []TestCaseForString{
		{Value: "ABC", Expected: "abc"},
		{Value: "aEiOu", Expected: "aeiou"},
		{Value: "AeIoU", Expected: "aeiou"},
		{Value: "aeiou", Expected: "aeiou"},
	} {
		testName := fmt.Sprintf("ToLowerCase(%v) === %v", tc.Expected, tc.Value)
		t.Run(testName, func(t2 *testing.T) {
			result := ToLowerCase(tc.Value)
			ExpectEqual(t2, result, tc.Expected)
		})
	}
}

func TestTrim(t *testing.T) {
	for _, tc := range []TestCaseForString{
		{Value: "    ", Expected: ""},
		{Value: "\t\t\t\t", Expected: ""},
		{Value: "\t\r\n\f", Expected: ""},
		{Value: " ABC ", Expected: "ABC"},
		{Value: "aEiOu", Expected: "aEiOu"},
		{Value: "   AeIoU", Expected: "AeIoU"},
		{Value: "AeIoU   ", Expected: "AeIoU"},
	} {
		testName := fmt.Sprintf("Trim(%v) === %v", tc.Expected, tc.Value)
		t.Run(testName, func(t2 *testing.T) {
			result := Trim(tc.Value)
			ExpectEqual(t2, result, tc.Expected)
		})
	}
}
