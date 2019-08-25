package bytes

import "math"

func SubSequences(bStr []byte) [][]byte {
	listLen := uint(len(bStr))
	subSeqLen := uint(math.Pow(2.0, float64(listLen)))
	out := make([][]byte, 0)
	var (
		i uint
		j uint
	)
	for i = 0; i < subSeqLen; i += 1 {
		entry := make([]byte, 0)
		for j = 0; j < listLen; j += 1 {
			if i & (1 << j) > 0 {
				entry = append(entry, bStr[int(j)])
			}
		}
		out = append(out, entry)
	}
	return out
}


func SliceSubSequences(bSlice [][]byte) [][][]byte {
	listLen := uint(len(bSlice))
	subSeqLen := uint(math.Pow(2.0, float64(listLen)))
	out := make([][][]byte, 0)
	var (
		i uint
		j uint
	)
	for i = 0; i < subSeqLen; i += 1 {
		entry := make([][]byte, 0)
		for j = 0; j < listLen; j += 1 {
			if i&(1<<j) > 0 {
				entry = append(entry, bSlice[int(j)])
			}
		}
		out = append(out, entry)
	}
	return out
}
