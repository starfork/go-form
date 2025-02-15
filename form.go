package goform

func NewVirtualForm(formStructData []byte) *VirtualForm {
	return &VirtualForm{}
}

func (e *VirtualForm) Validate() error {
	return nil
}

func (e *VirtualForm) Submit(data []byte) error {
	return nil
}

// func (f *Field) Parse() (any, error) {
// 	var parsed any
// 	var err error
// 	switch f.T {
// 	case "int":
// 		parsed, err = strconv.Atoi(val)
// 	case "float":
// 		parsed, err = strconv.ParseFloat(val, 64)
// 	case "string":
// 		parsed = val
// 	default:
// 		err = fmt.Errorf("未知类型: %s", field.Type)
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("字段 %s 解析错误: %v", field.Name, err)
// 	}
// 	return parsed, nil
// }
