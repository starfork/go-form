package goform

import (
	"fmt"
	"testing"
)

func TestGatewayToGameConfig(t *testing.T) {
	type TestConfig struct {
		V1 string
		V2 string
		V3 string
	}
	tc := &TestConfig{}
	data := []byte(`[{"k":"v1","v":"v1value"},{"k":"v2","v":"v2value"}]`)
	Struct(data, tc)
	fmt.Println(tc)
}
