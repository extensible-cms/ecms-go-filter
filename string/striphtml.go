package string

import (
	"bytes"
	"errors"
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"regexp"
)

// @todo copy implementation from fjl-filter
// https://github.com/functional-jslib/fjl-filter/blob/master/src/StripTagsFilter.js

var (
	blankSep = []byte("")

	// Character allowed for start of tag or html attribute name
	nameStartCharHexRanges = [][]byte{
		[]byte("\\x{C0}-\\x{D6}"),
		[]byte("\\x{D8}-\\x{F6}"),
		[]byte("\\x{F8}-\\x{2FF}"),
		[]byte("\\x{370}-\\x{37D}"),
		[]byte("\\x{37F}-\\x{1FFF}"),
		[]byte("\\x{200C}-\\x{200D}"),
		[]byte("\\x{2070}-\\x{218F}"),
		[]byte("\\x{2C00}-\\x{2FEF}"),
		[]byte("\\x{3001}-\\x{D7FF}"),
		[]byte("\\x{F900}-\\x{FDCF}"),
		[]byte("\\x{FDF0}-\\x{FFFD}"),
		[]byte("\\x{10000}-\\x{EFFFF}"),
	}

	// Characters allowed in tag names and/or html attribute names
	nameCharHexRanges = [][]byte{
		[]byte("\\x{0300}-\\x{036F}"),
		[]byte("\\x{203F}-\\x{2040}"),
	}

	nameStartCharPartial = bytes.Join(
		append(nameStartCharHexRanges, []byte(":_a-zA-Z")),
		[]byte(""),
	)

	nameCharPartial = bytes.Join(
		append(
			[][]byte{
				nameStartCharPartial,
				[]byte("\\-\\.0-9\\x{B7}"),
			},
			nameCharHexRanges...,
		),
		blankSep,
	)

	namePartial = bytes.Join([][]byte{
		[]byte("["), nameStartCharPartial, []byte("]"),
		[]byte("["), nameCharPartial, []byte("]*"),
	}, blankSep)

	eqPartial = []byte("\\s?=\\s?")

	multiLineSpacePartial = []byte("[\\n\\r\\t\\s]*")

	attrValuePartial = []byte("[^(?\\\")]*")

	attrPartial = bytes.Join([][]byte{namePartial, eqPartial, attrValuePartial}, blankSep)

	commentPartial = bytes.Join(
		[][]byte{
			[]byte("<!--"),
			multiLineSpacePartial,
			[]byte("(?m:.+)"),
			multiLineSpacePartial,
			[]byte("-->"),
		},
		blankSep,
	)

	commentRegex = regexp.MustCompile(string(commentPartial))

	tagNameRegex = regexp.MustCompile(string(namePartial))

	InvalidXmlTagNameError = errors.New("invalid xml tag name")

	InvalidAttribNameError = errors.New("invalid xml attribute name")
) // var declarations

func validateName(xs []byte) bool {
	return tagNameRegex.Match(xs)
}

func validateNames(xss [][]byte) (bool, [][]byte) {
	invalidNames := make([][]byte, 0)
	for _, xs := range xss {
		if validateName(xs) == false {
			invalidNames = append(invalidNames, xs)
		}
	}
	if len(invalidNames) > 0 {
		return false, invalidNames
	}
	return true, nil
}

func createTagRegexPartial(tagName []byte) []byte {
	return bytes.Join(
		[][]byte{
			[]byte("(<\\/?("), tagName, []byte(")(?:"),
			multiLineSpacePartial, attrPartial,
			[]byte(")*"),
			multiLineSpacePartial,
			[]byte(">)*"),
		},
		blankSep,
	)
}

func createAttribRegexPartial(attribName []byte) []byte {
	return bytes.Join(
		[][]byte{
			multiLineSpacePartial,
			[]byte("("),
			attribName,
			[]byte(")"),
		},
		blankSep,
	)
}

func GetStripHtmlTagsFilter(tagNames [][]byte) ecmsGoFilter.Filter {
	namesAreValid, _ /*invalidNames*/ := validateNames(tagNames)
	if !namesAreValid {
		// @todo add invalid tag names to error
		panic(InvalidXmlTagNameError)
	}
	if len(tagNames) == 0 {
		return ecmsGoFilter.Identity
	}
	return func(x interface{}) interface{} {
		bs := ecmsGoFilter.ToByteString(x)
		if bs == nil {
			return x
		}
		out := bs
		for _, tn := range tagNames {
			regex := regexp.MustCompile(string(createTagRegexPartial(tn)))
			out = regex.ReplaceAll(out, blankSep)
		}
		return out
	}
}

func GetStripHtmlAttribsFilter(attribNames [][]byte) ecmsGoFilter.Filter {
	namesAreValid, _ /*invalidNames*/ := validateNames(attribNames)
	if !namesAreValid {
		// @todo add invalid attrib names to error
		panic(InvalidAttribNameError)
	}
	if len(attribNames) == 0 {
		return ecmsGoFilter.Identity
	}
	return func(x interface{}) interface{} {
		bs := ecmsGoFilter.ToByteString(x)
		if bs == nil {
			return x
		}
		out := bs
		for _, tn := range attribNames {
			regex := regexp.MustCompile(string(createAttribRegexPartial(tn)))
			out = regex.ReplaceAll(out, blankSep)
		}
		return out
	}
}

func StripHtmlComments(x interface{}) interface{} {
	bs := ecmsGoFilter.ToByteString(x)
	if bs == nil {
		return x
	}
	return commentRegex.ReplaceAll(bs, blankSep)
}
