package redis

import "github.com/caarlos0/env/v9"

type PubConfig struct {
	LoggerDebug bool `env:"REDIS_PUB_SUB_LOGGER_DEBUG,notEmpty" envdefault:"false"`
	LoggerTrace bool `env:"REDIS_PUB_SUB_LOGGER_TRACE,notEmpty" envdefault:"false"`
}

func NewPubConfig() (*PubConfig, error) {
	c := &PubConfig{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
