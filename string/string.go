package string

import (
	"strings"
)

func LowerCase(xs interface{}) interface{} {
	return strings.ToLower(xs.(string))
}

func Trim(xs interface{}) interface{} {
	return strings.TrimSpace(xs.(string))
}

var (
	xmlEntitiesCharMap = map[byte][]byte{
		'<':  []byte("&lt;"),
		'>':  []byte("&gt;"),
		'"':  []byte("&quot;"),
		'\'': []byte("&apos;"),
		'&':  []byte("&amp;"),
	}
)

func XmlEntities(x interface{}) interface{} {
	var bs []byte
	switch x.(type) {
	case string:
		bs = []byte(x.(string))
	case []byte:
		bs = x.([]byte)
	case byte:
		bs = []byte{x.(byte)}
	default:
		return x
	}
	var out []byte
	for _, c := range bs {
		if xmlEntitiesCharMap[c] == nil {
			out = append(out, c)
			continue
		}
		out = append(out, xmlEntitiesCharMap[c]...)
	}
	return string(out)
}
