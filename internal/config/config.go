package config

import (
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"os"
	"strconv"
)

type Config struct {
	Log *zerolog.Logger
}

func NewConfig(i *do.Injector) (*Config, error) {
	return &Config{
		Log: do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

func (c *Config) GetInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	c.Log.Fatal().Err(err).Msg("Failed to parse int")

	return value
}
