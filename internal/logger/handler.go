package logger

import (
	"context"
	"log/slog"
	"os"

	publicAuth "github.com/atom-service/account/package/auth"
)

type Handler struct {
	slog.Handler
}

func NewHandler(level slog.Level) *Handler {
	return &Handler{
		Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		}),
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
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
