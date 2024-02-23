package auth

import (
	"context"
)

type TokenAuth struct {
	PublicKey string
	PrivateKey string
}

func (x *TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
  return map[string]string{
    "token":   x.Token,
  }, nil
}

func (x *TokenAuth) RequireTransportSecurity() bool {
  return false
}
