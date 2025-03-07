package config

type Logger struct {
	AppName     string `mapstructure:"app_name"`
	Environment string `mapstructure:"environment"`
	LogLevel    string `mapstructure:"log_level"`
}
