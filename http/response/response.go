package response

import (
	uuid "github.com/satori/go.uuid"
	"github.com/kenSevLeb/go-framework/component/paginate"
	"github.com/kenSevLeb/go-framework/component/trace"
)

const (
	// 请求ID前缀
	requestPrefix = "request:"
	// 成功提示信息
	successMessage = "success"
)

// Data 数据对象
type Data map[string]interface{}

// Response 响应对象
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	// 数据
	Data interface{} `json:"data"`
	// 元数据
	Meta Meta `json:"meta"`
}

// NewResponse 返回Response对象
func NewResponse() Response {
	return Response{
		Meta: Meta{
			RequestId: requestPrefix + uuid.NewV4().String(),
			TraceId:   trace.Get(),
		},
	}
}

// Meta 元数据
type Meta struct {
	// 请求ID
	RequestId string `json:"request_id"`
	// 全局追踪ID,服务化必备
	TraceId string `json:"trace_id"`
	// 分页数据项
	Pagination *paginate.Pagination `json:"pagination,omitempty"`
}
