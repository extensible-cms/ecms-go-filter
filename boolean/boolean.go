package boolean

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"github.com/extensible-cms/ecms-go-validator/is"
	"strings"
)

// @todo copy from fjl-filter implementation
// https://github.com/functional-jslib/fjl-filter/blob/master/src/BooleanFilter.js

type BoolFilterOptions struct {
	AllowCasting    bool
	Translations    map[string]string
	ConversionRules []string
}

func NewOptions() *BoolFilterOptions {
	return &BoolFilterOptions{
		AllowCasting:    true,
		Translations:    nil,
		ConversionRules: nil,
	}
}

func New(ops *BoolFilterOptions) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		switch x.(type) {
		case bool:
			return x
		}
		if !ops.AllowCasting {
			return x
		}
		if ops == nil {
			return castByNative(x)
		}
		if len(ops.ConversionRules) > 0 {
			return castByValue(x, ops)
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

func castByValue(x interface{}, ops *BoolFilterOptions) interface{} {
	if len(ops.ConversionRules) == 0 {
		return castByNative(x)
	}
	ruleFunctions := getConversionRules(ops.ConversionRules)
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
		var bs []byte
		switch x.(type) {
		case string:
			bs = []byte(x.(string))
		case []byte:
			bs = x.([]byte)
		default:
			return nil
		}
		lowerCasedStr := strings.ToLower(string(bs))
		for k, _ := range translations {
			if strings.ToLower(k) == lowerCasedStr {
				return translations[k]
			}
		}
		return false
	}
}

var (
	defaultTranslation = map[string]bool{
		"yes": true,
		"no": false,
		"undefined": false,
		"nil": false,
		"null": false,
		"0": false,
		"0.0": false,
	}
	castByTranslations = GetTranslationsCaster(defaultTranslation)
	castingRules = map[string]ecmsGoFilter.Filter{
		"native":       castByNative, 		// handles casting all value types
		"translations": castByTranslations, // default translations caster
	}
	ToBool ecmsGoFilter.Filter = New(NewOptions())
)
