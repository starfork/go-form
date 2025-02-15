package goform

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestExampleBuildForm(t *testing.T) {
	//formStructData := []byte(`{"action":"post","fields":[{"tt":"ApiKey","nm":"key","t":"string","r":1},{"tt":"ApiUrl","nm":"url","t":"string","r":1},{"tt":"Amt","nm":"amt","t":"float","r":1}]}`)
	formStructData2 := []byte(`{"action":"post","method":"post","ctxType":"","fields":[{"tt":"ApiKey","nm":"key","t":"string","r":1},{"tt":"ApiUrl","nm":"url","t":"string","r":1},{"tt":"Amt","nm":"amt","t":"float","r":-1}],"data":{}}`)
	f := &VirtualForm{}
	jsoniter.Unmarshal(formStructData2, f)
	//	fmt.Println(f)
	fmt.Println(len(f.Fields))
	for _, v := range f.Fields {

		fmt.Println(v.Nm)
	}

	testData2 := []byte(`[{"k":"key","v":"12345"},{"k":"url","v":"http://baidu.com"},{"k":"amt","v":"10"}]`)

	var params []Instance
	if err := jsoniter.Unmarshal(testData2, &params); err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Data)

	data := map[string]any{}
	for _, v := range params {
		for _, vv := range f.Fields {
			if v.K == vv.Nm {
				data[vv.Nm] = v.V
			}
		}
	}
	fmt.Println(data)
	// fmt.Println(form)
	// testData2 := []byte(`[{"k":"url","v":"http://baidu.com"},{"k":"v2","v":"v2value"}]`)
	// fmt.Println(testData2)
}
