// validator 参数校验器
// 通过标签的方式指定参数的校验方式，虽然性能有所损耗，但是能减少繁琐的参数判断
// 对程序代码的可读性有极大的提升
// Usage:
//
//	 	type User struct {
//				Name string `json:"phone" binding:"required,omitempty" comment:"名称"`
//				Age  uint   `json:"age" binding:"required,min=10" comment:"年龄"`
//			}
//
// 更多校验方式请查看:https://godoc.org/github.com/go-playground/validator
package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/kenSevLeb/go-framework/component/i18n"
	"reflect"
	"sync"
)

const (
	// 指定校验规则的tag
	validateTag = "binding"

	// 制定字段名称的注释,如 comment:"用户名"
	fieldNameTag = "comment"
)

func New() *Validator {
	v := new(Validator)

	return v
}

// Validator 自定义校验器
type Validator struct {
	once     sync.Once
	validate *validator.Validate
	locale   string
}

// getKindOf return the kind of data
func getKindOf(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// ValidateStruct validate struct
func (v *Validator) ValidateStruct(obj interface{}) error {

	if getKindOf(obj) == reflect.Struct {
		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {
			//验证器
			//for _, err := range err.(validator.ValidationErrors) {
			//	return errors.New(err.Translate(trans))
			//}
			return err

		}
	}

	return nil
}

// Engine
func (v *Validator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

// lazyInit
func (v *Validator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName(validateTag)

		// define filed name
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get(fieldNameTag)
		})

		// use translate
		i18n.RegisterTranslations(v.validate)

	})
}
