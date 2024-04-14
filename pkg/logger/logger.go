package logger

import (
	"log"
	"log/slog"
	"os"
)

type Logger struct {
	Logger *slog.Logger
}

func SetupLogger(modeLog string) *Logger {
	var handler slog.Handler

	switch modeLog {
	case "debug":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonDebug":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonInfo":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		log.Fatal("not init modelog: ", modeLog)
	}

	logger := &Logger{
		Logger: slog.New(handler),
	}

	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
