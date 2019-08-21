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
	testCases := make([]testCaseForBoolean, 0)
	for _, x := range []interface{}{"", 0, false, []interface{}{}} {
		testCases = append(testCases, testCaseForBoolean{
			Value:    x,
			Expected: false,
		})
	}

	for _, x := range []interface{}{"true", 1, true} {
		testCases = append(testCases, testCaseForBoolean{
			Value:    x,
			Expected: true,
		})
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("ToBool(%v) === `%v`", tc.Value, tc.Expected)
		t.Run(testName, func(t2 *testing.T) {
			ExpectEqual(t2, ToBool(tc.Value), tc.Expected)
		})
	}
}
