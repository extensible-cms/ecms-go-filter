package ecms_go_filter

import "math"

func ToByteString(x interface{}) []byte {
	switch x.(type) {
	case string:
		return []byte(x.(string))
	case []byte:
		return x.([]byte)
	case []rune:
		return []byte(string(x.([]rune)))
	case byte:
		return []byte{x.(byte)}
	case rune:
		return []byte(string(x.(rune)))
	default:
		return nil
	}
}

func StrSubSequences(xs string) []string {
	listLen := uint(len(xs))
	subSeqLen := uint(math.Pow(2.0, float64(listLen)))
	out := make([]string, 0)
	var (
		i uint
		j uint
	)
	for i = 0; i < subSeqLen; i += 1 {
		entry := make([]byte, 0)
		for j = 0; j < listLen; j += 1 {
			if i & (1 << j) > 0 {
				entry = append(entry, xs[int(j)])
			}
		}
		out = append(out, string(entry))
	}
	return out
}

func StrSliceSubSequences(xss []string) [][]string {
	listLen := uint(len(xss))
	subSeqLen := uint(math.Pow(2.0, float64(listLen)))
	out := make([][]string, 0)
	var (
		i uint
		j uint
	)
	for i = 0; i < subSeqLen; i += 1 {
		entry := make([]string, 0)
		for j = 0; j < listLen; j += 1 {
			if i&(1<<j) > 0 {
				entry = append(entry, xss[int(j)])
			}
		}
		out = append(out, entry)
	}
	return out
}
