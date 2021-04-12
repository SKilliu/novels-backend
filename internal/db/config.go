package db

import "github.com/caarlos0/env"

var configuration Configuration

type Configuration struct {
	Name     string `env:"db_name,required"`
	Host     string `env:"db_host,required"`
	Port     int    `env:"db_port,required"`
	User     string `env:"db_user,required"`
	Password string `env:"db_password,required"`
	SSL      string `env:"db_ssl"`
}

func loadConfigFromEnvs() {

	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to get db data from env")
		panic(err)
	}
}
