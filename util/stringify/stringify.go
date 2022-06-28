package stringify

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type String interface {
	String() string
}

type Int interface {
	Int() int64
}

type Uint interface {
	Uint() uint64
}

type Float interface {
	Float() float64
}

type Bool interface {
	Bool() bool
}

func ToString(i interface{}, subs ...interface{}) string {

	valueElement := GetReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid {
		return ""
	}

	if e, ok := valueElement.Interface().(String); ok {
		return e.String()
	}

	if valueTypeKind == reflect.String {
		return valueElement.String()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() == true {
			return "true"
		}
		return "false"
	}

	if valueTypeKind <= reflect.Int64 {
		return strconv.FormatInt(valueElement.Int(), 10)
	}

	if valueTypeKind <= reflect.Uintptr {
		return strconv.FormatUint(valueElement.Uint(), 10)
	}

	if valueTypeKind <= reflect.Float64 {
		if len(subs) > 0 {
			f := ToInt(subs[0])
			return strconv.FormatFloat(valueElement.Float(), 'f', int(f), 64)
		}
		return fmt.Sprintf("%v", valueElement.Float())
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			return string(e)
		}

		sli := []string{}
		sub := ","
		subSubs := subs
		if len(subs) > 0 {
			sub = ToString(subs[0])
			subSubs = subSubs[1:]
		}

		for i := 0; i < valueElement.Len(); i++ {
			value := ToString(valueElement.Index(i), subSubs...)
			sli = append(sli, value)
		}

		return strings.Join(sli, sub)
	}

	return ""
}

func GetReflectValue(v interface{}) reflect.Value {
	if e, ok := v.(reflect.Value); ok {
		return getElem(e)
	}

	return getElem(reflect.ValueOf(v))
}

func getElem(v reflect.Value) reflect.Value {

	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return getElem(v.Elem())
	}

	return v
}

func ToInt(i interface{}) int64 {

	valueElement := GetReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid {
		return 0
	}

	if e, ok := valueElement.Interface().(Int); ok {
		return e.Int()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() == true {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return valueElement.Int()
	}

	if valueTypeKind <= reflect.Uintptr {
		return int64(valueElement.Uint())
	}

	if valueTypeKind <= reflect.Float64 {
		return int64(valueElement.Float())
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseInt(valueElement.String(), 10, 64)
		return v
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseInt(string(e), 10, 64)
			return v
		}

		return int64(valueElement.Len())
	}

	return 0
}

func ToUint(i interface{}) uint64 {

	valueElement := GetReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid {
		return 0
	}

	if e, ok := valueElement.Interface().(Uint); ok {
		return e.Uint()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() == true {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return uint64(valueElement.Int())
	}

	if valueTypeKind <= reflect.Uintptr {
		return valueElement.Uint()
	}

	if valueTypeKind <= reflect.Float64 {
		return uint64(valueElement.Float())
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseUint(valueElement.String(), 10, 64)
		return v
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseUint(string(e), 10, 64)
			return v
		}

		return uint64(valueElement.Len())
	}

	return 0
}

func ToFloat(i interface{}) float64 {

	valueElement := GetReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid {
		return 0
	}

	if e, ok := valueElement.Interface().(Float); ok {
		return e.Float()
	}

	if valueTypeKind == reflect.Bool {
		if valueElement.Bool() == true {
			return 1
		}
		return 0
	}

	if valueTypeKind <= reflect.Int64 {
		return float64(valueElement.Int())
	}

	if valueTypeKind <= reflect.Uintptr {
		return float64(valueElement.Uint())
	}

	if valueTypeKind <= reflect.Float64 {
		return valueElement.Float()
	}

	if valueTypeKind == reflect.String {
		v, _ := strconv.ParseFloat(valueElement.String(), 10)
		return v
	}

	if valueTypeKind == reflect.Slice {

		if e, ok := valueElement.Interface().([]byte); ok {
			v, _ := strconv.ParseFloat(string(e), 10)
			return v
		}

		return float64(valueElement.Len())
	}

	return 0
}

func ToBool(i interface{}, subs ...interface{}) bool {

	valueElement := GetReflectValue(i)
	valueTypeKind := valueElement.Kind()

	if valueTypeKind == reflect.Invalid {
		return false
	}

	if e, ok := valueElement.Interface().(Bool); ok {
		return e.Bool()
	}

	if valueTypeKind == reflect.Bool {
		return valueElement.Bool()
	}

	if valueTypeKind <= reflect.Int64 {
		if valueElement.Int() != 0 {
			return true
		}
		return false
	}

	if valueTypeKind <= reflect.Uintptr {
		if valueElement.Uint() != 0 {
			return true
		}
		return false
	}

	if valueTypeKind <= reflect.Float64 {
		if valueElement.Float() != 0 {
			return true
		}
		return false
	}

	if valueTypeKind == reflect.String {
		str := valueElement.String()

		if len(subs) > 0 {
			f := ToBool(subs[0])
			if f == true {
				if str == "" || str == "0" || str == "false" {
					return false
				}
				return true
			}
		}

		if str != "" {
			return true
		}
	}

	if valueTypeKind == reflect.Slice {
		if e, ok := valueElement.Interface().([]byte); ok {
			return ToBool(string(e), subs...)
		}
	}

	return false
}

func ToStringSlice(i string, subs ...interface{}) []string {

	if i == "" {
		return []string{}
	}

	sli := []string{i}
	sub := []string{}
	if len(subs) > 0 {
		subElement := GetReflectValue(subs[0])
		if subElement.Kind() == reflect.Slice {
			for i := 0; i < subElement.Len(); i++ {
				sub = append(sub, ToString(subElement.Index(i)))
			}
		}
		if subElement.Kind() == reflect.String {
			sub = append(sub, subElement.String())
		}
	} else {
		sub = append(sub, ",")
	}

	for _, v := range sub {
		s := []string{}
		for _, v2 := range sli {
			s = append(s, strings.Split(v2, v)...)
		}
		sli = s
	}

	return sli

}

func ToIntSlice(i string, subs ...interface{}) []int64 {

	stringSlice := ToStringSlice(i, subs...)
	intSlice := make([]int64, len(stringSlice))

	for k, v := range stringSlice {
		intSlice[k], _ = strconv.ParseInt(v, 10, 64)
	}

	return intSlice
}
