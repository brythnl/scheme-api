package logger

import (
	"io"
	"log/slog"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(w, nil)),
	}
}
