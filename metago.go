package metago

import (
	"reflect"
	"fmt"
)

func CallFunc(f interface{}, methodName string, args ...interface{}) interface{} {
	var ptr reflect.Value
	var value reflect.Value

	value, ok := f.(reflect.Value)

	if !ok {
		value = reflect.ValueOf(f)
	}

	funcArgs := convertFuncArgs(args)

	ptr, value = getReflectPtrAndValue(value, f)

	method := value.MethodByName(methodName)

	if method.IsValid() {
		return unwrapFuncResult(method.Call(funcArgs))
	}

	method = ptr.MethodByName(methodName)

	if method.IsValid() {
		return unwrapFuncResult(method.Call(funcArgs))
	}

	panic(fmt.Sprintf("%T type objects don't have %s method.", value.Interface(), methodName))
}

func convertFuncArgs(args []interface{}) []reflect.Value {
	funcArgs := []reflect.Value{}

	for _, arg := range args {
		value, wrapped := arg.(reflect.Value)

		if wrapped {
			funcArgs = append(funcArgs, value)
			continue
		}

		funcArgs = append(funcArgs, reflect.ValueOf(arg))
	}

	return funcArgs
}


func unwrapFuncResult(result interface{}) interface{} {
	switch result := result.(type) {
	case []reflect.Value:
		if len(result) == 0 {
			return nil
		} else if len(result) == 1 {
			return result[0].Interface()
		} else {
			values := []interface{}{}

			for _, v := range result {
				values = append(values, v.Interface())
			}

			return values
		}
	default:
		return result
	}
}

func getReflectPtrAndValue(value reflect.Value, rawValue interface{}) (ptr, v reflect.Value) {
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(rawValue)) // create new pointer
		temp := ptr.Elem()                          // create variable to value of pointer
		temp.Set(value)                             // set value of variable to our passed in value
	}

	return ptr, value
}