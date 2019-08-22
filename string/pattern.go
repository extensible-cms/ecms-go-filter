package string

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"regexp"
)

// GetPatternFilter returns a filter that filters strings and byte-strings using pattern
//  and replacement passed in.
func GetPatternFilter(pattern *regexp.Regexp, replacement []byte) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		bs := ecmsGoFilter.ToByteString(x)
		if bs == nil {
			return x
		}
		return pattern.ReplaceAll(bs, replacement)
	}
}

var (
	slugRegex       = regexp.MustCompile("[^a-z\\-_\\d]")
	slugReplacement = []byte("-")

	// Slug filter takes a string or byte string and filters it using given pattern
	Slug ecmsGoFilter.Filter = GetPatternFilter(slugRegex, slugReplacement)
)
