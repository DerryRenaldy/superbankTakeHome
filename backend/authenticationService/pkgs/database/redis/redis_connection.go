package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string `json:"address"  mapstructure:"address"`
	Timeout  int    `json:"timeout" mapstructure:"timeout"`
	PoolSize int    `json:"pool_size" mapstructure:"pool_size"`
}

type TokenCache struct {
	AccessTokenTimeout  int `json:"access_token_timeout" mapstructure:"access_token_timeout"`
	RefreshTokenTimeout int `json:"refresh_token_timeout" mapstructure:"refresh_token_timeout"`
}

func (c *Config) Parse() (*redis.Options, error) {
	opts, err := redis.ParseURL(c.Addr)
	if err != nil {
		return nil, err
	}

	opts.PoolSize = c.PoolSize
	if opts.PoolSize <= 0 {
		opts.PoolSize = 1000
	}

	return opts, err
}

// ConnectRedis ...
func ConnectRedis(cfg *Config) (*redis.Client, error) {
	opts1, err := cfg.Parse()
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts1)
	if err = client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}