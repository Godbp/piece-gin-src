package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InfoWithContext 上下文日志记录
func InfoWithContext(ctx *gin.Context, format string, opts ...interface{}) {
	params := getCtxParamsLog(ctx)
	fields := logrus.Fields{
		"reqId": params.ReqID,
	}
	log.WithFields(fields).Infof(format, opts...)
}
