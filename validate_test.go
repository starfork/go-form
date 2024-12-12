package goform

import (
	"fmt"
	"testing"
)

func TestValidateTemplate(t *testing.T) {
	pm := `[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`
	fmt.Println(ValidateTemplate(pm))
}

func TestValidateInstance(t *testing.T) {
	pm := `[{"k":"api","v":"apiapi"},{"k":"key","v":"keykey"}]`
	tplPm := `[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`
	fmt.Println(ValidateInstance(pm, tplPm))
}
