package logger

import (
	"context"
	"log/slog"

	"github.com/atom-service/account/internal/config"
)

func Init(ctx context.Context) error {
	var level = slog.LevelInfo

	if config.Logger != nil {
		if config.Logger.Level != "" {
			if config.Logger.Level == "debug" {
				level = slog.LevelDebug
			}

			if config.Logger.Level == "info" {
				level = slog.LevelInfo
			}

			if config.Logger.Level == "warn" {
				level = slog.LevelWarn
			}

			if config.Logger.Level == "error" {
				level = slog.LevelError
			}
		}
	}

	handler:= NewHandler(level)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}
