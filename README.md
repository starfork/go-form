# go-form

一个 Go 动态表单处理库，支持**表单模板验证**、**表单实例验证**和**动态结构体转换**。验证器通过接口抽象，支持自定义实现。

## 安装

```bash
go get github.com/starfork/go-form
```

如需使用 [go-playground/validator](https://github.com/go-playground/validator) 作为验证引擎，额外引入适配子包：

```bash
go get github.com/starfork/go-form/goplay
```

## 核心类型

### Template（表单模板）

```json
[
  {"nm": "name", "tt": "姓名", "t": "input", "r": "required"},
  {"nm": "age",  "tt": "年龄", "t": "input", "r": "gte=0,lte=150"},
  {"nm": "bio",  "tt": "简介", "t": "textarea"}
]
```

| 字段 | 含义 | 说明 |
|------|------|------|
| `nm` | 字段名 | 必填 |
| `tt` | 字段标题 | 必填 |
| `t`  | 表单类型 | 可选：`input` / `select` / `textarea` |
| `r`  | 验证规则 | 验证器使用的规则字符串（如 `required`, `gte=0`） |

### Instance（表单实例）

```json
[
  {"k": "name", "v": "张三"},
  {"k": "age",  "v": 25}
]
```

## 验证

验证器通过 `Validator` 接口抽象：

```go
type Validator interface {
    Struct(s any) error
    Var(field any, rule string) error
}
```

核心包 **不内置任何验证器实现**，你可以选择：
- 使用 `goplay` 子包（基于 go-playground/validator）
- 实现自己的验证器

### 使用 go-playground

```go
import (
    goform "github.com/starfork/go-form"
    "github.com/starfork/go-form/goplay"
)

func main() {
    v := goplay.New()

    // 验证模板
    tpl := []byte(`[{"nm":"api","tt":"名称"},{"nm":"key","tt":"密钥"}]`)
    goform.ValidateTemplate(tpl, v)

    // 验证实例
    pm := []byte(`[{"k":"api","v":11},{"k":"key","v":"keykey"}]`)
    tplPm := []byte(`[{"nm":"api","tt":"名称","r":"gt=10"},{"nm":"key","tt":"密钥"}]`)
    goform.ValidateInstance(pm, tplPm, v)
}
```

### 使用自定义验证器

实现 `Validator` 接口即可接入任意验证库：

```go
import goform "github.com/starfork/go-form"

type SimpleValidator struct{}

func (SimpleValidator) Struct(s any) error {
    // 自定义校验逻辑
    return nil
}

func (SimpleValidator) Var(field any, rule string) error {
    // 自定义校验逻辑
    return nil
}

func main() {
    v := SimpleValidator{}
    goform.ValidateTemplate([]byte(`[{"nm":"x","tt":"X"}]`), v)
}
```

### 不传验证器（基础校验）

不传验证器时，库会自动进行基础校验（仅检查必填字段非空）：

```go
goform.ValidateTemplate([]byte(`[{"nm":"api","tt":"名称"}]`)) // OK
goform.ValidateInstance(pm, tpl) // 仅检查 key 是否存在，跳过规则验证
```

## 动态结构体转换

将 `[]Instance` JSON 按字段名匹配（首字母大写）动态赋值到任意 Go 结构体：

```go
import goform "github.com/starfork/go-form"

type User struct {
    Name string
    Age  int64
    Bio  string
}

data := []byte(`[{"k":"name","v":"张三"},{"k":"age","v":25}]`)

u := &User{}
goform.Struct(data, u) // &User{Name:"张三", Age:25}
```

支持自动类型转换：`string` / `int*` / `float*` / `bool`。

## API

### 验证

```go
func Validate(pm, tpl []byte, validate ...Validator) error
func ValidateTemplate(pm []byte, validate ...Validator) error
func ValidateInstance(pm, tpl []byte, validate ...Validator) error
```

- `Validate` — 自动判断：有模板时校验实例，无模板时校验模板
- `ValidateTemplate` — 校验模板参数完整性
- `ValidateInstance` — 按模板规则校验实例数据

### 转换

```go
func Struct(jsonData []byte, target any) error
```
