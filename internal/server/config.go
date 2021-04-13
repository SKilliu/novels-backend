package server

import (
	"github.com/caarlos0/env"
)

type Configuration struct {
	Host           string `env:"server_host"`
	Port           int    `env:"server_port"`
	SSL            bool   `env:"server_ssl"`
	ServerCertPath string `env:"server_cert_path"`
	ServerKeyPath  string `env:"server_key_path"`
}

var configuration Configuration

func loadConfigFromEnvs() {

	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to get server config from env")
		panic(err)
	}
}
