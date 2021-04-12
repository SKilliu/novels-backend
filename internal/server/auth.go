package server

import "github.com/caarlos0/env"

var authConfig Auth

type Auth struct {
	VerifyKey string `env:"authentication_secret,required"`
	Algorithm string `env:"authentication_algorithm"`
}

func loadAuthConfigFromEnvs() {
	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to get auth config from env")
		panic(err)
	}
}
