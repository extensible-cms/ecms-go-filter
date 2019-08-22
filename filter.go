package ecms_go_filter

type Filter = func(interface{}) interface{}

func Identity(x interface{}) interface{} {
	return x
}
