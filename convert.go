package goform

import (
	"fmt"
	"reflect"
	"strconv"

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
				switch fieldVal.Kind() {
				case reflect.String:
					fieldVal.SetString(param.V)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if intVal, err := strconv.ParseInt(param.V, 10, 64); err == nil {
						fieldVal.SetInt(intVal)
					}
				case reflect.Float32, reflect.Float64:
					if floatVal, err := strconv.ParseFloat(param.V, 64); err == nil {
						fieldVal.SetFloat(floatVal)
					}
				case reflect.Bool:
					if boolVal, err := strconv.ParseBool(param.V); err == nil {
						fieldVal.SetBool(boolVal)
					}
				default:
					fmt.Printf("Unsupported field type: %s\n", fieldVal.Kind())
				}
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
