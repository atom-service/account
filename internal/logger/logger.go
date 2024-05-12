package logger

import (
	"context"
	"log/slog"

	"github.com/atom-service/account/internal/config"
)

func Init(ctx context.Context) error {
	if config.Logger != nil {
		if config.Logger.Level != "" {
			if config.Logger.Level == "debug" {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}

			if config.Logger.Level == "info" {
				slog.SetLogLoggerLevel(slog.LevelInfo)
			}

			if config.Logger.Level == "warn" {
				slog.SetLogLoggerLevel(slog.LevelWarn)
			}

			if config.Logger.Level == "error" {
				slog.SetLogLoggerLevel(slog.LevelError)
			}
		}
	}

	return nil
}
