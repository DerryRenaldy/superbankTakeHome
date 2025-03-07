package authcachestore

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthCache struct {
	client           *redis.Client
	namespace        string
	timeoutInSeconds int
}

func (c AuthCache) GetToken(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.timeoutInSeconds)*time.Second)
	defer cancel()

	return c.client.Get(ctx, c.redisKey(key)).Result()
}

func (c AuthCache) SetToken(ctx context.Context, key string, value string, ttl int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.timeoutInSeconds)*time.Second)
	defer cancel()

	return c.client.Set(ctx, c.redisKey(key), value, time.Duration(ttl)*time.Minute).Err()
}

func (c AuthCache) DeleteToken(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.timeoutInSeconds)*time.Second)
	defer cancel()

	return c.client.Del(ctx, c.redisKey(key)).Err()
}

func (c AuthCache) redisKey(key string) string {
	return fmt.Sprintf("%v:%v", c.namespace, key)
}

func New(client *redis.Client, namespace string, timeoutInSeconds int) *AuthCache {
	return &AuthCache{
		client:           client,
		namespace:        namespace,
		timeoutInSeconds: timeoutInSeconds,
	}
}
