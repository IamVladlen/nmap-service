package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Log struct {
	zerolog.Logger
}

// New creates new logger with specified level.
func New(logLevel string) *Log {
	l := zerolog.New(os.Stdout)
	logger := l.With().Timestamp().Logger().Level(setLevel(logLevel))

	return &Log{
		logger,
	}
}

func setLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		fmt.Printf("Incorrect log level: edit /config/config.yaml\n")
		return zerolog.InfoLevel
	}
}
