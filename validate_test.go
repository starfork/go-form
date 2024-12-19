package goform

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidateTemplate(t *testing.T) {
	pm := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateTemplate(pm))
}
func TestValidate(t *testing.T) {
	pm := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(Validate(pm, nil))
}

func TestValidateInstance(t *testing.T) {
	pm := []byte(`[{"k":"api","v":11},{"k":"key","v":"keykey"}]`)
	tplPm := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateInstance(pm, tplPm))
	tplPm2 := []byte(`[{"nm":"api","tt":"名称","r":"gt=10"},{"nm":"key","tt":"密钥"}]`)
	fmt.Println(ValidateInstance(pm, tplPm2))
}

func TestValidateVar(t *testing.T) {
	v := validator.New()
	var tv any = 123
	fmt.Println(v.Var(tv, "gt=123"))
	tv = "ss"
	fmt.Println(v.Var(tv, "gt=1"))
}
