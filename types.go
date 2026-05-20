package goform

// Validator 验证器接口，支持自定义实现
type Validator interface {
	Struct(s any) error
	Var(field any, rule string) error
}

// 表单模版
type Template struct {
	Tt string `validate:"required" json:"tt"` //字段标题
	Nm string `validate:"required" json:"nm"` //字段名
	//表单类型
	T string `validate:"omitempty,oneof=input select textarea" json:"t"`
	//表单验证规则
	R string `json:"r"`
}

type Instance struct {
	K string `json:"k"` //字段名
	V any    `json:"v"` //字段值
}


 
