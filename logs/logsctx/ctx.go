package logsctx

import (
	"context"

	"github.com/uber-go/zap"
	"os"
	"path"
)

const requestIdKey = 1

var logger zap.Logger

func init() {
	// a fallback/root logger for events without context
	logger = zap.New(
		zap.NewJSONEncoder(zap.RFC3339Formatter("key")),
		zap.Fields(zap.Int("pid", os.Getpid()),
			zap.String("exe", path.Base(os.Args[0]))),
	)
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
	}
	return newLogger
}
