package log

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func Infof(format string, v ...interface{}) {
	InfoContextf(nil, format, v...)
}

func InfoContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Infof(ctx, format, v...)
}

func Debugf(format string, v ...interface{}) {
	DebugContextf(nil, format, v...)
}

func DebugContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Debugf(ctx, format, v...)
}

func Errorf(format string, v ...interface{}) {
	ErrorContextf(nil, format, v...)
}

func ErrorContextf(ctx context.Context, format string, v ...interface{}) {
	g.Log().Errorf(ctx, format, v...)
}
