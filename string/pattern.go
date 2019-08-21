package string

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"regexp"
)

// GetPatternFilter returns a filter that filters strings and byte-strings using pattern
//  and replacement passed in.
func GetPatternFilter(pattern *regexp.Regexp, replacement []byte) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		var bs []byte
		switch x.(type) {
		case string:
			bs = []byte(x.(string))
		case []byte:
			bs = x.([]byte)
		default:
			return x
		}
		return pattern.ReplaceAll(bs, replacement)
	}
}

var (
	slugRegex = regexp.MustCompile("[^a-z\\-_\\d]")
	slugReplacement = []byte("-")

	// Slug filter takes a string or byte string and filters it using given pattern
	Slug ecmsGoFilter.Filter = GetPatternFilter(slugRegex, slugReplacement)
)
