package util

import (
	"context"

	ctxTagsInterceptor "github.com/ffauzann/common/interceptor/grpc/unary/ctxtags"
	"go.uber.org/zap"
)

var log *zap.Logger

func SetLogger(logger *zap.Logger) {
	log = logger
}

func Log() *zap.Logger {
	return log
}

func LogContext(ctx context.Context) *zap.Logger {
	tags := ctxTagsInterceptor.Extract(ctx)
	values := tags.Values()
	cID, _ := values[ctxTagsInterceptor.CIDKey].(string)
	return log.With(zap.String(ctxTagsInterceptor.CIDKey, cID))
}
