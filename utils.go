package ecms_go_filter

func ToByteString(x interface{}) []byte {
	switch x.(type) {
	case string:
		return []byte(x.(string))
	case []byte:
		return x.([]byte)
	case byte:
		return []byte{x.(byte)}
	default:
		return nil
	}
}
