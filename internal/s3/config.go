package s3

import "github.com/caarlos0/env"

var configuration Configuration

type Configuration struct {
	AccessKey string `env:"s3_access_key"`
	SecretKey string `env:"s3_secret_key"`
	Url       string `env:"s3_url"`
	Bucket    string `env:"s3_bucket"`
}

func loadConfigFromEnvs() {

	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to load s3 config from envs")
		panic(err)
	}
}
