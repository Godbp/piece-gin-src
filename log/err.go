package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log 应用程序日志句柄
var log *logrus.Logger

// ErrorWithContext 上下文错误日志记录
func ErrorWithContext(ctx *gin.Context, format string, opts ...interface{}) {
	params := getCtxParamsLog(ctx)
	fields := logrus.Fields{
		"reqId": params.ReqID,
	}
	log.WithFields(fields).Errorf(format, opts...)
}
