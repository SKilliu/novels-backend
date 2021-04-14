package server

import (
	"github.com/caarlos0/env"
)

var authConfig Auth

type Auth struct {
	VerifyKey     string `env:"authentication_secret,required"`
	EmailAddress  string `env:"email_address"`
	EmailPassword string `env:"email_password"`
}

func loadAuthConfigFromEnvs() {
	if err := env.Parse(&authConfig); err != nil {
		logger.WithError(err).Error("failed to get auth config from env")
		panic(err)
	}
}
