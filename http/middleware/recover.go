package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kenSevLeb/go-framework/component/i18n"
	"github.com/kenSevLeb/go-framework/component/log"
	"github.com/kenSevLeb/go-framework/errors"
	"github.com/kenSevLeb/go-framework/http/response"
	"github.com/kenSevLeb/go-framework/util/strings"
	"runtime"
)

// 处理因系统panic导致的错误，需要拦截，定制化返回响应内容，不能直接输出与系统相关的内容，可以自定义
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			log.Error("Recover", log.Content{"param": r, "uri": ctx.Request.RequestURI})

			switch err := r.(type) {
			case validator.ValidationErrors:
				response.Wrap(ctx).Error(1001, i18n.TranslateValidatorErrors(i18n.LANG_ZH, err))
			case *errors.Error:
				response.Wrap(ctx).Error(err.Code, err.Message)
			case error: // 普通error
				fmt.Printf("panic: %v\n", r)
				printStack() // 只要当panic的非validateError和自定义Error才打印stack
				response.Wrap(ctx).Error(500, err.Error())
			default:
				fmt.Printf("panic: %v\n", r)
				printStack() // 只要当panic的非validateError和自定义Error才打印stack
				response.Wrap(ctx).Error(500, "system error")
			}
		}
	}()
	ctx.Next()
}

// 多语言
func I18nRecover(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

				log.Error("Recover", log.Content{"param": r, "uri": ctx.Request.RequestURI})

				// 获取当前语言，默认为zh
				lang := strings.Default(ctx.GetString("lang"), i18n.LANG_ZH)
				wrapper := response.Wrap(ctx)
				switch err := r.(type) {
				case validator.ValidationErrors:
					wrapper.Error(1001, i18n.TranslateValidatorErrors(lang, err))
				case *errors.Error:
					response.Wrap(ctx).Error(err.Code, bundle.MustLocalize(lang, err.Error()))
				case error: // 普通error
					fmt.Printf("panic: %v\n", r)
					printStack() // 只要当panic的非validateError和自定义Error才打印stack
					response.Wrap(ctx).Error(500, err.Error())
				default:
					fmt.Printf("panic: %v\n", r)
					printStack() // 只要当panic的非validateError和自定义Error才打印stack
					response.Wrap(ctx).Error(500, "system error")
				}
			}
		}()
		ctx.Next()
	}

}

func printStack() {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}
