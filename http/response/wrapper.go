package response

import (
	"github.com/gin-gonic/gin"
	"github.com/kenSevLeb/go-framework/component/paginate"
	"reflect"
)

// Wrapper context装饰器
type Wrapper struct {
	ctx *gin.Context
}

// Wrap
func Wrap(ctx *gin.Context) *Wrapper {
	return &Wrapper{
		ctx: ctx,
	}
}

// output output response
func (w *Wrapper) output(response Response) {
	w.ctx.JSON(200, response)
	w.ctx.Abort()
}

// Success 输出成功响应
func (w *Wrapper) Success(data interface{}) {
	response := NewResponse()
	response.Message = successMessage
	response.Data = data

	w.output(response)
}

// Error 输出错误响应
func (w *Wrapper) Error(code int, message string) {
	response := NewResponse()
	response.Code = code
	response.Message = message

	w.output(response)
}

// Output 自定义输出响应
func (w *Wrapper) Output(code int, message string, data interface{}) {
	response := NewResponse()
	response.Data = data
	response.Code = code
	response.Message = message

	w.output(response)
}

// Paginate 输出分页响应内容
func (w *Wrapper) Paginate(data interface{}, pagination *paginate.Pagination) {
	response := NewResponse()
	// 如果data为nil,则默认设置为[]
	if data == nil || reflect.ValueOf(data).IsNil() {
		data = []interface{}{}
	}

	response.Data = data
	response.Meta.Pagination = pagination

	w.output(response)
}
