package config

import (
	"authenticationService/pkgs/database/redis"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret string        `mapstructure:"jwt_secret"`
	App       Logger        `mapstructure:",squash"`
	DB        PostgresDatabase `mapstructure:",squash"`
	Redis     redis.Config  `mapstructure:"redis"`
	TokenCache redis.TokenCache `mapstructure:"token_cache"`
}

var Cfg *Config

func init() {
	// Check if running tests
	if os.Getenv("GO_TEST") != "" {
		return
	}

	fmt.Println("Reading config file...")

	Cfg = loadConfig()
	fmt.Println(Cfg)
}

func loadConfig() *Config {
	cfg := new(Config)

	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AddConfigPath("/root/config")
	err := v.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(fmt.Errorf("fatal error config file not found: %w", err))
		} else {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	if errUnmarshal := v.Unmarshal(cfg); errUnmarshal != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))

	}

	return cfg
}
