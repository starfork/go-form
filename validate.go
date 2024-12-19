package goform

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

// 自动验证
func Validate(pm, tpl []byte, validate ...*validator.Validate) error {
	if len(tpl) == 0 {
		return ValidateTemplate(pm, validate...)
	}
	return ValidateInstance(pm, tpl, validate...)
}

// 验证模版
func ValidateTemplate(pm []byte, validate ...*validator.Validate) error {
	var tpls []Template
	if err := jsoniter.Unmarshal(pm, &tpls); err != nil {
		return fmt.Errorf("template data json error %+v", err)
	}
	var v *validator.Validate
	if len(validate) > 0 {
		v = validate[0]
	} else {
		v = validator.New()
	}

	for i, item := range tpls {
		if err := v.Struct(item); err != nil {
			return fmt.Errorf("第 %d 组参数不完整: 参数需要以 nm,tt 成对出现且不能空", i+1)
		}
	}

	return nil
}

// 验证实例
func ValidateInstance(pm, tpl []byte, validate ...*validator.Validate) error {
	var tpls []Template
	if err := jsoniter.Unmarshal(tpl, &tpls); err != nil {
		return fmt.Errorf("template data json error %+v", err)
	}
	var instances []Instance
	if err := jsoniter.Unmarshal(pm, &instances); err != nil {
		return fmt.Errorf("instance data json error %+v", err)
	}

	var v *validator.Validate
	if len(validate) > 0 {
		v = validate[0]
	} else {
		v = validator.New()
	}

	tplMap := make(map[string]string)
	for _, tpl := range tpls {
		tplMap[tpl.Nm] = tpl.R
	}
	for i, instance := range instances {
		rule, exists := tplMap[instance.K]
		if !exists {
			return fmt.Errorf("第 %d: 键 '%s' 不在模版参数中", i+1, instance.K)
		}

		if rule != "" {

			if err := v.Var(instance.V, rule); err != nil {
				return fmt.Errorf("第 %d: 键  '%s' 校验失败, 值 '%s' 不符合规则 '%s'", i+1, instance.K, instance.V, rule)
			}
		}
	}
	return nil
}
