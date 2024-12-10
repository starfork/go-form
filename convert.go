package goform

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

func Struct(jsonData []byte, target interface{}) error {

	var params []Instance
	if err := jsoniter.Unmarshal(jsonData, &params); err != nil {
		return err
	}
	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	for _, param := range params {
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.Name == "" || field.Name != capitalize(param.K) {
				continue
			}
			fieldVal := val.Field(i)
			if fieldVal.CanSet() {
				fieldVal.SetString(param.V)
			}
		}
	}
	return nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[0]-32) + s[1:]
}
