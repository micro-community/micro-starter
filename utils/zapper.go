package utils

import (
	"go.uber.org/zap"
)

// 性能高，用于中间件
var ZLogger *zap.Logger

// 普通打印
var Sugar *zap.SugaredLogger

func init() {
	ZLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Sugar = ZLogger.Sugar()
}
