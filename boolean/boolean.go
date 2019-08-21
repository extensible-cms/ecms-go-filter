package boolean

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
)

// @todo copy from fjl-filter implementation
// https://github.com/functional-jslib/fjl-filter/blob/master/src/BooleanFilter.js

type BoolFilterOptions struct {
	AllowCasting    bool
	Translations    map[string]string
	ConversionRules []string
}

func NewBoolFilterOptions() *BoolFilterOptions {
	return &BoolFilterOptions{
		AllowCasting:    true,
		Translations:    nil,
		ConversionRules: nil,
	}
}

func NewBoolFilter(ops *BoolFilterOptions) ecmsGoFilter.Filter {
	return func(x interface{}) interface{} {
		if !ops.AllowCasting {
			return x
		}
		conversionRulesLen := len(ops.ConversionRules)
		switch x.(type) {
		case bool:
			return x
		case string:
			if conversionRulesLen > 0 {
				return castValue(x, ops)
			}
		}
		if ops.AllowCasting && conversionRulesLen == 0 &&
			len(ops.Translations) == 0 {
			return castByNative(x)
		}
		return false
	}
}

func castByNative(x interface{}) bool {
	return false
}

var (
	castingRules = map[string]string{
		"byNative": "castByNative",
	}
	ToBool ecmsGoFilter.Filter = NewBoolFilter(NewBoolFilterOptions())
)

func castValue(x interface{}, ops *BoolFilterOptions) bool {
	return false
}
