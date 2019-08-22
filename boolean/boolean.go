package boolean

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"github.com/extensible-cms/ecms-go-validator/is"
	"strings"
)

// @note Prior art: fjl-filter
// https://github.com/functional-jslib/fjl-filter/blob/master/src/BooleanFilter.js

func GetBoolFilter(allowCasting bool, conversionRules []string) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		switch x.(type) {
		case bool:
			return x
		}
		if !allowCasting {
			return x
		}
		if len(conversionRules) > 0 {
			return castByValue(x, conversionRules)
		}
		return castByNative(x)
	}
}

// castByNative Takes care of casting any go primitive value to a boolean
// E.g., ```
//   for _, x := range []interface{}{[]string{}{}, nil, struct{}{}, 0, false, ""} {
//     castByNative(x) == false // statement is `true` for all empty values and vice versa
//	 }
// ```
func castByNative(x interface{}) interface{} {
	return !is.Empty(x)
}

func castByValue(x interface{}, conversionRules []string) interface{} {
	if len(conversionRules) == 0 {
		return castByNative(x)
	}
	ruleFunctions := getConversionRules(conversionRules)
	for _, f := range ruleFunctions {
		result := f(x)
		switch result.(type) {
		case bool:
			return result.(bool)
		}
	}
	return false
}

func getConversionRules(conversionRules []string) []ecmsGoFilter.Filter {
	out := make([]ecmsGoFilter.Filter, 0)
	lenConversionRules := len(conversionRules)

	if lenConversionRules == 0 {
		return out
	}

	// If conversion rules includes only "all" key, get all conversion functions
	if lenConversionRules == 1 && strings.ToLower(conversionRules[0]) == "all" {
		for _, fn := range castingRules {
			out = append(out, fn)
		}
		return out
	}

	// Else get only requested conversion functions
	for _, k := range conversionRules {
		rule := strings.ToLower(k)
		if castingRules[rule] != nil {
			out = append(out, castingRules[rule])
		}
	}

	// Return functions
	return out
}

func GetTranslationsCaster(translations map[string]bool) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		bs := ecmsGoFilter.ToByteString(x)
		if bs == nil {
			return nil
		}
		lowerCasedStr := strings.ToLower(string(bs))
		for k := range translations {
			if strings.ToLower(k) == lowerCasedStr {
				return translations[k]
			}
		}
		return nil
	}
}

var (
	defaultTranslation = map[string]bool{
		"":          false,
		"yes":       true,
		"no":        false,
		"true":      true,
		"false":     false,
		"undefined": false,
		"nil":       false,
		"null":      false,
		"1":         true,
		"0":         false,
		"1.0":       true,
		"0.0":       false,
		"{}":        false,
		"[]":        false,
	}
	castByTranslations = GetTranslationsCaster(defaultTranslation)
	castingRules       = map[string]ecmsGoFilter.Filter{
		"native":       castByNative,       // handles casting all value types
		"translations": castByTranslations, // default translations caster
	}
	ToBool = GetBoolFilter(true, nil)
)
