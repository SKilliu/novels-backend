package server

import (
	"os"
)

var authConfig Auth

type Auth struct {
	VerifyKey string `env:"authentication_secret,required"`
	Algorithm string `env:"authentication_algorithm"`
}

func loadAuthConfigFromEnvs() {
	secret, ok := os.LookupEnv("authentication_secret")
	if !ok {
		secret = "12312312323"
	}
	authConfig.VerifyKey = secret

	algorithm, ok := os.LookupEnv("authentication_algorithm")
	if !ok {
		algorithm = "HS256"
	}
	authConfig.Algorithm = algorithm
}
