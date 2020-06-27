package reflectionbits

import "reflect"

func getValue(x interface{}) reflect.Value{
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr{
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, f func(input string)){
	// Returns value initialized stored in interface x
	val := getValue(x)


	walkValue := func(value reflect.Value) {
		walk(value.Interface(), f)
	}

	switch val.Kind() {
	case reflect.String:
		f(val.String())
	case reflect.Struct:
		for i := 0; i< val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i:= 0; i<val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v,ok = val.Recv(){
			walk(v.Interface(), f)
		}
	case reflect.Func:
		valFnRsult := val.Call(nil)
		for _, res := range valFnRsult{
			walk(res.Interface(), f)
		}
	}

}

