package boolean

import (
	"fmt"
	"testing"
)

func ExpectEqual(tx *testing.T, a interface{}, b interface{}) {
	if a != b {
		tx.Errorf("Expected: %v;  Received: %v", b, a)
	}
}

func TestToBool(t *testing.T) {
	type testCaseForBoolean struct {
		Value    interface{}
		Expected bool
	}

	// Gather test cases here
	testCases := make([]testCaseForBoolean, 0)

	// Falsy test cases
	for _, x := range []interface{}{"", 0, false, []interface{}{}} {
		testCases = append(testCases, testCaseForBoolean{
			Value:    x,
			Expected: false,
		})
	}

	// Truthy test cases
	for _, x := range []interface{}{"true", 1, true} {
		testCases = append(testCases, testCaseForBoolean{
			Value:    x,
			Expected: true,
		})
	}

	// Run test cases
	for _, tc := range testCases {
		testName := fmt.Sprintf("ToBool(%v) === `%v`", tc.Value, tc.Expected)
		t.Run(testName, func(t2 *testing.T) {
			ExpectEqual(t2, ToBool(tc.Value), tc.Expected)
		})
	}
}
