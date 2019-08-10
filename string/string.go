package string

import "strings"

func ToLowerCase(xs interface{}) interface{} {
	return strings.ToLower(xs.(string))
}

func Trim(xs interface{}) interface{} {
	return strings.TrimSpace(xs.(string))
}
