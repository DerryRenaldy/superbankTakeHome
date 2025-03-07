package authcacherepo

import (
	authcache "authenticationService/stores/auth_cache"
	"context"
)

type IRepository interface {
	GetToken(ctx context.Context, key string) (string, error)
	SetToken(ctx context.Context, key string, value string, ttl int) error
	DeleteToken(ctx context.Context, key string) error
}

var _ IRepository = (*authcache.AuthCache)(nil)