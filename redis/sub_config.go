package redis

import "github.com/caarlos0/env/v9"

type SubConfig struct {
	PubConfig
	ConsumerGroup string `env:"REDIS_PUB_SUB_CONSUMER_GROUP_ID,notEmpty" envdefault:""`
}

func NewSubConfig() (*SubConfig, error) {
	c := &SubConfig{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
