package server

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type Configuration struct {
	Host           string
	Port           int
	SSL            bool
	ServerCertPath string
	ServerKeyPath  string
}

var configuration Configuration

func loadConfigFromEnvs() {
	var err error

	port, ok := os.LookupEnv("server_port")
	if !ok {
		port = "8081"
	}
	configuration.Port, err = strconv.Atoi(port)
	if err != nil {
		panic(errors.Wrap(err, "failed to convert port to int"))
	}

	host, ok := os.LookupEnv("server_host")
	if !ok {
		host = "localhost"
	}
	configuration.Host = host

	ssl, ok := os.LookupEnv("server_ssl")
	if !ok || ssl != "true" {
		configuration.SSL = false
	} else {
		configuration.SSL = true
	}

	serverCertPath, ok := os.LookupEnv("server_cert_path")
	if !ok {
		configuration.ServerCertPath = "ssl/server_cert_path"
	}
	configuration.ServerCertPath = serverCertPath

	serverKeyPath, ok := os.LookupEnv("server_key_path")
	if !ok {
		configuration.ServerKeyPath = "ssl/server_key_path"
	}
	configuration.ServerKeyPath = serverKeyPath
}
