package email

import "github.com/caarlos0/env"

var configuration Configuration

type Configuration struct {
	SendgridAPIKey string `env:"SENDGRID_API_KEY"`
	Email          string `env:"email"`
	Author         string `env:"email_author"`
}

func loadAuthConfigFromEnvs() {
	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to get server config from env")
		panic(err)
	}
}
