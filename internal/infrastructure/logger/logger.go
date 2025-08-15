package logger

import (
	"fmt"
	"io"
	"log/slog"
)

func New(w io.Writer) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(w, nil))
	fmt.Printf("Logger created: %p\n", logger)

	return logger
}
