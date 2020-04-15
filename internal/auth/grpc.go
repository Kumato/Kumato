package auth

import "context"

func (jc JWTClaims) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"INTERNAL_TOKEN": InternalToken(),
	}, nil
}

func (jc JWTClaims) RequireTransportSecurity() bool {
	return true
}
