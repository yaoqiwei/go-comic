package convert

import (
	"encoding/json"
	"reflect"
	"time"

	"fehu/util/stringify"
)

func IsType(val interface{}, t string) bool {
	return t == reflect.TypeOf(val).Name()
}

func StructAssign(binding interface{}, value interface{}) {

	bindingElement := reflect.ValueOf(binding).Elem()
	valueElement := reflect.ValueOf(value).Elem()

	valueKeyNum := valueElement.NumField()

	for i := 0; i < valueKeyNum; i++ {
		valueField := valueElement.Field(i)
		valueFieldName := valueElement.Type().Field(i).Name
		bindingField := bindingElement.FieldByName(valueFieldName)

		if ok := bindingField.IsValid(); !ok {
			continue
		}
		if ok := bindingField.CanSet(); !ok {
			continue
		}

		if ok := valueField.IsZero(); ok {
			continue
		}

		valueFieldValue := reflect.ValueOf(valueField.Interface())

		if bindingField.Kind() == valueFieldValue.Kind() && valueFieldValue.Type().ConvertibleTo(bindingField.Type()) {
			bindingField.Set(valueFieldValue.Convert(bindingField.Type()))
			continue
		}

		if valueFieldValue.Type().Name() == "Time" && bindingField.Kind() == reflect.String {
			if time, ok := valueFieldValue.Interface().(time.Time); ok {
				bindingField.SetString(time.Format("2006-01-02 15:04:05"))
				continue
			}
		}

		bindingTypeField, _ := bindingElement.Type().FieldByName(valueFieldName)
		sub := bindingTypeField.Tag.Get("sub")
		sub2 := valueElement.Type().Field(i).Tag.Get("sub")

		TypeChange(bindingField, valueFieldValue, sub, sub2)
	}
}

// value的值赋值给binding
// 支持bool,intx,uintx,string,floatx互相转换
// 支持[]string <=> string转换
func TypeChange(binding interface{}, value interface{}, subs ...interface{}) {

	bindingElement := stringify.GetReflectValue(binding)
	valueElement := stringify.GetReflectValue(value)

	bindingTypeKind := bindingElement.Type().Kind()
	valueTypeKind := valueElement.Type().Kind()

	if bindingTypeKind >= reflect.Int && bindingTypeKind <= reflect.Int64 {
		bindingElement.SetInt(stringify.ToInt(valueElement))
	} else if bindingTypeKind == reflect.String {
		bindingElement.SetString(stringify.ToString(valueElement, subs...))
	} else if bindingTypeKind == reflect.Bool {
		bindingElement.SetBool(stringify.ToBool(valueElement, subs...))
	} else if bindingTypeKind <= reflect.Uintptr {
		bindingElement.SetUint(stringify.ToUint(valueElement))
	} else if bindingTypeKind <= reflect.Float64 {
		bindingElement.SetFloat(stringify.ToFloat(valueElement))
	} else if bindingTypeKind == reflect.Slice {
		if _, ok := bindingElement.Interface().([]string); !ok {
			return
		}
		if valueTypeKind == reflect.String {
			b := stringify.ToStringSlice(valueElement.String())
			bindingElement.Set(reflect.ValueOf(b))
		}
	}

}

func InArrayString(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

func JoinMap(s map[string]interface{}, i ...[]string) ([]string, []interface{}) {

	var in []string
	var ex []string
	var key []string
	var val []interface{}

	if len(i) > 0 {
		in = i[0]
	}
	if len(i) > 1 {
		ex = i[0]
	}

	for k, v := range s {
		if ex != nil && InArrayString(k, ex) {
			continue
		}

		if in != nil && !InArrayString(k, in) {
			continue
		}
		key = append(key, k)
		val = append(val, v)
	}

	return key, val

}

func ToJson(i interface{}, e ...string) string {
	b, err := json.Marshal(i)
	if err != nil {
		if len(e) > 0 {
			return e[0]
		}
		return ""
	}
	return string(b)
}
