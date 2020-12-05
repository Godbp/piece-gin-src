package log

import "github.com/gin-gonic/gin"

// reqParam 请求参数
type reqParam struct {
	ReqID string `json:"reqId"` // 请求跟踪号 uuid4
}

// getCtxParamsLog 获取日志参数信息
func getCtxParamsLog(ctx *gin.Context) *reqParam {
	return &reqParam{
		ReqID: ctx.Writer.Header().Get("X-Request-Id"),
	}
}
