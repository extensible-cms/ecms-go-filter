package string

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
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
	bs := ecmsGoFilter.ToByteString(x)
	if bs == nil {
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
