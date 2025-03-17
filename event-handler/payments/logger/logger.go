package logger

import (
	"context"
	"event-handler/payments/config"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

type Logger struct {
	*slog.Logger
	level slog.Level
}

var Log Logger

const (
	CorrelationIDKey = echo.HeaderXRequestID
)

func (l *Logger) WithError(err error) *Logger {
	n := l.With(slog.Attr{
		Key:   "error",
		Value: slog.AnyValue(err),
	})

	return &Logger{n, l.level}
}

func (l *Logger) WithAny(k string, val any) *Logger {
	n := l.With(slog.Attr{
		Key:   k,
		Value: slog.AnyValue(val),
	})

	return &Logger{n, l.level}
}

func New(s *config.Config) *Logger {
	level := slog.Level(s.LogLevel)

	opts := &slog.HandlerOptions{Level: level}
	base := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	Log = Logger{base, level}

	return &Log
}

func FromCtx(ctx context.Context) Logger {
	val := ctx.Value(CorrelationIDKey)
	newLog := *Log.WithAny(CorrelationIDKey, val)
	return newLog
}
