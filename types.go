package goform

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

// 表单字段。这个相当于显示的生命字段内容类型
// 一般使用在动态api提交的时候(http请求的api请求，不需要特别严格的参数类型的场景)
// 实际上对于url.Values(params).Encode() 是不需要类型的。
// 所以大多数情况是为了适配“application/json”的场景
type Field struct {
	Tt string `json:"tt"` //字段标题
	Nm string `json:"nm"` //字段名
	//表单数据类型。一般的只需要区分api类型的大致类型即可。
	//表单类型,int,string,float
	T string `json:"t"` //

	R int `json:"r"` //是否必填。1必须
}

type VirtualForm struct {
	Action  string `json:"action"` //提交地址
	Method  string `json:"method"` //一般都是post
	CtxType string `json:"ctxType"`

	Fields []*Field       `json:"fields"` //表单字段
	Data   map[string]any `json:"data"`   //表单数据
}
