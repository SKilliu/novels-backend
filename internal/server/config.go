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

	// var err error

	// port, ok := os.LookupEnv("server_port")
	// if !ok {
	// 	port = "8081"
	// }
	// configuration.Port, err = strconv.Atoi(port)
	// if err != nil {
	// 	panic(errors.Wrap(err, "failed to convert port to int"))
	// }

	// host, ok := os.LookupEnv("server_host")
	// if !ok {
	// 	host = "localhost"
	// }
	// configuration.Host = host

	// ssl, ok := os.LookupEnv("server_ssl")
	// if !ok || ssl != "true" {
	// 	configuration.SSL = false
	// } else {
	// 	configuration.SSL = true
	// }

	// serverCertPath, ok := os.LookupEnv("server_cert_path")
	// if !ok {
	// 	configuration.ServerCertPath = "ssl/server_cert_path"
	// }
	// configuration.ServerCertPath = serverCertPath

	// serverKeyPath, ok := os.LookupEnv("server_key_path")
	// if !ok {
	// 	configuration.ServerKeyPath = "ssl/server_key_path"
	// }
	// configuration.ServerKeyPath = serverKeyPath
}
