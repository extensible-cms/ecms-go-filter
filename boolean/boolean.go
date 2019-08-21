package boolean

import (
	ecmsGoFilter "github.com/extensible-cms/ecms-go-filter"
	"github.com/extensible-cms/ecms-go-validator/is"
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

var (
	castingRules = map[string]string{
		"byNative": "castByNative",
	}
	ToBool ecmsGoFilter.Filter = New(NewOptions())
)

func castByNative(x interface{}) bool {
	return !is.Empty(x)
}

func castValue(x interface{}, ops *BoolFilterOptions) bool {
	return false
}

// @note Every casting function from here down should return `interface{}`;  E.g.,
// 	`nil` if couldn't/didn't convert value else a boolean

func castBoolean(x interface{}) interface{} {
	switch x.(type) {
	case bool:
		return x.(bool)
	default:
		return nil
	}
}

func castInt(x interface{}) interface{} {
	switch x.(type) {
	case int:
		return x.(int) != 0
	case int64:
		return x.(int64) != 0
	case int32:
		return x.(int32) != 0
	default:
		return nil
	}
}

func castUint(x interface{}) interface{} {
	switch x.(type) {
	case uint64:
		return x.(uint64) != 0
	case uint32:
		return x.(uint32) != 0
	default:
		return nil
	}
}

func castFloat(x interface{}) interface{} {
	switch x.(type) {
	case float64:
		return x.(float64) != 0.0
	case float32:
		return x.(float32) != 0.0
	default:
		return nil
	}
}
