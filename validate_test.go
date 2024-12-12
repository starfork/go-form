package goform

import (
	"fmt"
	"testing"
)

func TestValidateTemplate(t *testing.T) {
	pm := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateTemplate(pm))
}

func TestValidateInstance(t *testing.T) {
	pm := []byte(`[{"k":"api","v":"apiapi"},{"k":"key","v":"keykey"}]`)
	tplPm := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateInstance(pm, tplPm))
	tplPm2 := []byte(`[{"nm":"api","tt":"名称","r":"gt=10"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateInstance(pm, tplPm2))
}
