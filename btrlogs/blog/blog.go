package blog

import (
	"context"
	"log/slog"
	"os"
)

type Key string

var RequestIDKey Key = "X-Request-ID"

func New(ctx context.Context) *slog.Logger {
	requestID := ctx.Value(RequestIDKey)
	if requestID == nil {
		requestID = "unknown"
	}

	h := slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}).WithAttrs(
		[]slog.Attr{
			slog.String("app", "btrlogs"),
			slog.String("request_id", requestID.(string)),
		})

	return slog.New(h)
}
