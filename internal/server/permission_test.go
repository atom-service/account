package server

import "context"

func withAdminPermission(ctx context.Context) context.Context {
	return ctx
}

func withTokenPermission(ctx context.Context, token string) context.Context {
	return ctx
}

func withUnknownPermission(ctx context.Context) context.Context {
	return ctx
}
