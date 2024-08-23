package logger

import (
	"log/slog"
	"os"
)


func NewLoger() (*slog.Logger, error) {
	opts := slog.HandlerOptions{
		Level:  slog.LevelDebug,
	}

	file, err := os.OpenFile("apilog.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(file, &opts))

	return logger, nil
}