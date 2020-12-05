package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// // WarningWithContext 上下文日志记录
func WarningWithContext(ctx *gin.Context, format string, opts ...interface{}) {
	params := getCtxParamsLog(ctx)
	fields := logrus.Fields{
		"reqId": params.ReqID,
	}
	log.WithFields(fields).Warnf(format, opts...)
}
