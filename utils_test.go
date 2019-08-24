package ecms_go_filter

import (
	"fmt"
	"math"
	"testing"
)

func TestStrSubSequences(t *testing.T) {
	for _, candidate := range []string{
		"",
		"abc",
	} {
		t.Run(fmt.Sprintf("StrSubSequences(\"%v\")", candidate), func(t2 *testing.T) {
			candidateLen := len(candidate)
			result := StrSubSequences(candidate)
			resultLen := len(result)
			expectedLen := int(math.Pow(2, float64(candidateLen)))

			t2.Run(fmt.Sprintf("Expect result length of %v", expectedLen), func(t2 *testing.T) {
				if resultLen != expectedLen {
					t2.Errorf("Expected %v;  Got %v", expectedLen, resultLen)
				}
			})

			// Ensure correct subsequences were generated
			for i, x := range result {
				name := fmt.Sprintf("Expect lastIndexOf %v to equal %v", x, i)
				t2.Run(name, func(t2 *testing.T) {
					for j, x2 := range result {
						if i != j && x == x2 {
							t2.Errorf("Expected last index of %v to equal %v;  Got %v", x, i, j)
						}
					}
				})
			}
		})
	}
}

func TestToByteString(t *testing.T) {
	for _, x := range []interface{}{
		"",
		' ',
		"abc",
		[]byte(""),
		[]byte{' '},
		[]byte("abc"),
		[]rune(""),
		[]rune{' '},
		[]rune("abc"),
	} {
		t.Run(fmt.Sprintf("ToByteString(%v)", x), func(t2 *testing.T) {
			result := ToByteString(x)
			resultStr := string(result)
			var expectedStr string
			switch x.(type) {
			case string:
				expectedStr = x.(string)
			case []byte:
				expectedStr = string(x.([]byte))
			case []rune:
				expectedStr = string(x.([]rune))
			case rune:
				expectedStr = string([]rune{x.(rune)})
			}

			t2.Run(fmt.Sprintf("Expect %v same as %v", result, expectedStr), func(t3 *testing.T) {
				if resultStr != expectedStr {
					t3.Errorf("Expected %v to %v", resultStr, expectedStr)
				}
			})
		})
	}
}
