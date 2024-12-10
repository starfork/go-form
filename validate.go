package goform

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

func ValidateTemplate(pm string, validate ...*validator.Validate) error {
	var tpls []Template
	if err := jsoniter.Unmarshal([]byte(pm), &tpls); err != nil {
		return err
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

func ValidateInstance(pm, tplPm string, validate ...*validator.Validate) error {
	var tpls []Template
	if err := jsoniter.Unmarshal([]byte(pm), &tpls); err != nil {
		return err
	}
	var v *validator.Validate
	if len(validate) > 0 {
		v = validate[0]
	} else {
		v = validator.New()
	}

	var instances []Instance
	if err := jsoniter.Unmarshal([]byte(pm), &instances); err != nil {
		return err
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
