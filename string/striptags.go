package string

import (
	"bytes"
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
)

// @todo copy implementation from fjl-filter
// https://github.com/functional-jslib/fjl-filter/blob/master/src/StripTagsFilter.js

var (
	rangeSep = []byte{'-'}

	blankSep = []byte("")

	// Character allowed for start of tag or html attribute name
	nameStartCharUnicodeRanges = [][]byte{
		[]byte("\\u{C0}-\\u{D6}"),
		[]byte("\\u{D8}-\\u{F6}"),
		[]byte("\\u{F8}-\\u{2FF}"),
		[]byte("\\u{370}-\\u{37D}"),
		[]byte("\\u{37F}-\\u{1FFF}"),
		[]byte("\\u{200C}-\\u{200D}"),
		[]byte("\\u{2070}-\\u{218F}"),
		[]byte("\\u{2C00}-\\u{2FEF}"),
		[]byte("\\u{3001}-\\u{D7FF}"),
		[]byte("\\u{F900}-\\u{FDCF}"),
		[]byte("\\u{FDF0}-\\u{FFFD}"),
		[]byte("\\u{10000}-\\u{EFFFF}"),
	}

	// Characters allowed in tag names and/or html attribute names
	nameCharUnicodeRanges = [][]byte{
		[]byte("\\u{0300}-\\u{036F}"),
		[]byte("\\u{203F}-\\u{2040}"),
	}

	nameStartCharPartial = bytes.Join(
		append(nameStartCharUnicodeRanges, []byte(":_a-zA-Z")),
		[]byte(""),
	)

	nameCharPartial = bytes.Join(
		append(
			[][]byte{
				nameStartCharPartial,
				[]byte("\\-\\.0-9\\u{B7}"),
			},
			nameCharUnicodeRanges...,
		),
		blankSep,
	)

	namePartial = bytes.Join([][]byte{
		[]byte("["), nameStartCharPartial, []byte("]"),
		[]byte("["), nameCharPartial, []byte("]*"),
	}, blankSep)

	eqPartial = []byte("\\s?=\\s?")

	multiLineSpacePartial = []byte("[\\n\\r\\t\\s]*")

	attrValuePartial = []byte("")

) // var declarations

func stripEscapeSeqHead(xs []byte) []byte {
	return xs[2:]
}

func wrapUnicodeClass(xs []byte) []byte {
	return []byte("\\u{" + string(xs) + "}")
}

func hexRangeToUnicodeRange(rangeStr []byte) []byte {
	parts := bytes.Split(rangeStr, rangeSep)
	fromDigit := stripEscapeSeqHead(parts[0])
	toDigit := stripEscapeSeqHead(parts[1])
	fromBs := wrapUnicodeClass(fromDigit)
	toBs := wrapUnicodeClass(toDigit)
	return bytes.Join([][]byte{fromBs, toBs}, rangeSep)
}

func GetStripTagsFilter() ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		return x
	}
}
