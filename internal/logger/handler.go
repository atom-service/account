package logger

import (
	"context"
	"log/slog"
	"os"

	publicAuth "github.com/atom-service/account/package/auth"
)

type handler struct {
	slog.Handler
}

func NewHandler(level slog.Level) *handler {
	return &handler{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: level.Level() == slog.LevelDebug,
			Level:     level,
		}),
	}
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	authData := publicAuth.ResolveAuth(ctx)

	if authData != nil {
		if authData.User != nil && authData.User.ID != nil {
			r.Add(slog.Int64("userID", *authData.User.ID))
		}

		if authData.Secret != nil && authData.Secret.Key != nil {
			r.Add(slog.String("secretKey", *authData.Secret.Key))
		}
	}

	return h.Handler.Handle(ctx, r)
}
