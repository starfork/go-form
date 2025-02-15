package goform

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type testStruct struct {
	V1 string
	V2 string
	V3 float64
	V4 uint32
}

var (
	testData  = []byte(`[{"k":"v1","v":"v1value"},{"k":"v2","v":"v2value"}]`)
	testData1 = []byte(`[{"k":"v3","v":"10"},{"k":"v2","v":"v2value"}]`)
)

func TestConvertToStruct(t *testing.T) {

	tc := &testStruct{}

	Struct(testData, tc)
	fmt.Println(tc)
	data2 := []Instance{
		{K: "v3", V: "32"},
	}
	b1, _ := json.Marshal(data2)
	Struct(b1, tc)
	fmt.Println(tc)

	Struct(testData1, tc)
	fmt.Println(tc)

}
func TestConvertStructVar(t *testing.T) {
	v := validator.New()
	tc := &testStruct{}
	Struct(testData1, tc)
	fmt.Println(v.Var(tc.V3, "gt=11"))
	fmt.Println(v.Var(tc.V3, "gte=10"))
}

//type exampleForm struct{}
