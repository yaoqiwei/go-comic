package validator

import (
	"fehu/model/http_error"
	"reflect"
	"regexp"
	"strings"
)

// 判断变量是否为空
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func IsMail(str string) bool {
	match, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", str)
	return match
}

func IsPhone(str string) bool {
	match, _ := regexp.MatchString("^1[3|4|5|6|7|8|9]\\d{9}$", str)
	return match
}

func IsUserName(str string) bool {
	match, _ := regexp.MatchString("^[0-9a-zA-Z]{3,20}$", str)
	return match
}

func Passcheck(str string) bool {
	// num, _ := regexp.MatchString("^[a-zA-Z]+$", str)
	// word, _ := regexp.MatchString("^[0-9]+$", str)
	check, _ := regexp.MatchString("^[a-zA-Z0-9]{6,12}$", str)

	// if num || word {
	// 	panic(http_error.PasswordTypeError)
	// }

	if !check {
		panic(http_error.PasswordCountError)
	}

	return true
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
