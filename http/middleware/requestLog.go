package middleware

import (
	"bytes"
	"git.yingxiong.com/platform/go-framework/component/log"
	"git.yingxiong.com/platform/go-framework/component/trace"
	"git.yingxiong.com/platform/go-framework/util/date"
	"git.yingxiong.com/platform/go-framework/util/strings"
	"github.com/gin-gonic/gin"
	"time"
)

// bodyLogWriter 读取响应Writer
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

const (
	maxResponseSize = 2000 // 最大响应长度
)

// Write 读取响应数据
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLog gin的请求日志中间件
func RequestLog(c *gin.Context) {
	t := time.Now()
	requestTime := date.GetLocalMicroTimeStampStr()
	blw := &bodyLogWriter{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
	c.Writer = blw

	requestBody := strings.GetRequestBody(c.Request)

	c.Next()

	// package log content
	items := log.Content{}
	items["request_uri"] = c.Request.RequestURI
	items["request_method"] = c.Request.Method
	items["refer_service_name"] = c.Request.Referer()
	items["refer_request_host"] = c.ClientIP()
	items["request_body"] = requestBody
	items["request_time"] = requestTime
	items["response_time"] = date.GetLocalMicroTimeStampStr()
	items["response_body"] = getResponseBody(blw.body.String())
	items["time_used"] = time.Since(t).Microseconds()
	//items["header"] = c.Request.Header

	// 使用异步的方式写入日志，提高并发性能
	trace.Go(func() {
		log.Info("REQUEST_LOG", items)
	})
}

// getResponseBody 获取响应内容
func getResponseBody(s string) string {
	if strLength := len(s); strLength > maxResponseSize {
		res := make([]byte, maxResponseSize)
		copy(res, s[:maxResponseSize])
		return string(res)
	}
	return s
}
