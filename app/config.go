package app

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	AppName          string `env:"APP_NAME"`
	AppPort          int    `env:"APP_PORT" envDefault:"3000"`
	LogLevel         string `env:"LOG_LEVEL"`
	Environment      string `env:"ENVIRONMENT"`
	DatabaseDsn      string `env:"DATABASE_DSN"`
	FinanceApiDomain string `env:"FINANCE_API_DOMAIN"`
}

func LoadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg
}
