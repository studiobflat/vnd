package postgre

import (
	"strings"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/jackc/pgx/v5/tracelog"
)

type LogLevel tracelog.LogLevel

func (t *LogLevel) UnmarshalText(text []byte) error {
	l, err := tracelog.LogLevelFromString(strings.ToLower(string(text)))
	if err != nil {
		*t = LogLevel(tracelog.LogLevelError)
	}
	*t = LogLevel(l)
	return nil
}

type Config struct {
	Url                   string        `env:"POSTGRE_URL,required"`
	MaxConnection         int32         `env:"POSTGRE_MAX_CONNECTION,required"`
	MinConnection         int32         `env:"POSTGRE_MIN_CONNECTION,required"`
	MaxConnectionIdleTime time.Duration `env:"POSTGRE_MAX_IDLE_TIME,required"`
	LogLevel              LogLevel      `env:"POSTGRE_LOG_LEVEL" envDefault:"ERROR"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
