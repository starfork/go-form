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
			if !fieldVal.CanSet() {
				continue
			}

			//value := reflect.ValueOf(param.V)
			switch fieldVal.Kind() {
			case reflect.String:
				if v, ok := param.V.(string); ok {
					fieldVal.SetString(v)
				} else {
					fmt.Printf("Field %d: Cannot set non-string value '%v' to string field\n", i, param.V)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intVal, err := toInt(param.V)
				if err == nil {
					fieldVal.SetInt(intVal)
				} else {
					fmt.Printf("Field %d: Cannot convert '%v' to int: %v\n", i, param.V, err)
				}
			case reflect.Float32, reflect.Float64:
				floatVal, err := toFloat(param.V)
				if err == nil {
					fieldVal.SetFloat(floatVal)
				} else {
					fmt.Printf("Field %d: Cannot convert '%v' to float: %v\n", i, param.V, err)
				}
			case reflect.Bool:
				boolVal, err := toBool(param.V)
				if err == nil {
					fieldVal.SetBool(boolVal)
				} else {
					fmt.Printf("Field %d: Cannot convert '%v' to bool: %v\n", i, param.V, err)
				}
			default:
				fmt.Printf("Field %d: Unsupported field type: %s\n", i, fieldVal.Kind())
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

func toInt(v any) (int64, error) {
	switch val := v.(type) {
	case string:
		return strconv.ParseInt(val, 10, 64)
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(val).Int(), nil
	case float32, float64:
		return int64(reflect.ValueOf(val).Float()), nil
	default:
		return 0, fmt.Errorf("unsupported type '%T'", v)
	}
}

func toFloat(v any) (float64, error) {
	switch val := v.(type) {
	case string:
		return strconv.ParseFloat(val, 64)
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(val).Int()), nil
	case float32, float64:
		return reflect.ValueOf(val).Float(), nil
	default:
		return 0, fmt.Errorf("unsupported type '%T'", v)
	}
}

func toBool(v any) (bool, error) {
	switch val := v.(type) {
	case string:
		return strconv.ParseBool(val)
	case bool:
		return val, nil
	default:
		return false, fmt.Errorf("unsupported type '%T'", v)
	}
}
