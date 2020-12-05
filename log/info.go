package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InfoWithContext(ctx *gin.Context, format string, opts ...interface{}) {
	params := getCtxParamsLog(ctx)
	fields := logrus.Fields{
		"reqId": params.ReqID,
	}
	log.WithFields(fields).Infof(format, opts...)
}
