package util

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidStruct = errors.New("invalid struct specified")
)

// IsEmpty returns true if value empty
func IsEmpty(v reflect.Value) bool {
	switch getKind(v) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int:
		return v.Int() == 0
	case reflect.Uint:
		return v.Uint() == 0
	case reflect.Float32:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return IsEmpty(v.Elem())
	case reflect.Func:
		return v.IsNil()
	case reflect.Invalid:
		return true
	case reflect.Struct:
		var ok bool
		for i := 0; i < v.NumField(); i++ {
			ok = IsEmpty(v.FieldByIndex([]int{i}))
			if !ok {
				return false
			}
		}
	default:
		return false
	}
	return true
}

// Zero creates new zero interface
func Zero(src interface{}) (interface{}, error) {
	sv := reflect.ValueOf(src)

	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}

	if sv.Kind() == reflect.Invalid {
		return nil, ErrInvalidStruct
	}

	dst := reflect.New(sv.Type())

	return dst.Interface(), nil
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()
	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	}
	return kind
}
