package logger

import (
	"context"
	"log/slog"
	"os"

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

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}))

	slog.SetDefault(logger)

	return nil
}
