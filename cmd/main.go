package main

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/SKilliu/novels-backend/internal/server"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// const pathToConfigFile = "./envs.yaml"

// @title Novels REST API
// @version 1.0.2
// @description REST API for Novels app.
// @description New in version:<br> - fixed bug with voting.
// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
// @host 165.227.207.77:8000
// @BasePath /
func main() {
	log := logrus.New()
	logger := logrus.NewEntry(log)

	// utils.UploadEnvironmentVariables(pathToConfigFile)

	db.Init(logger)

	server.Init(logger)

	err := server.Start()
	if err != nil {
		panic(errors.Wrap(err, "failed to start api"))
	}
}
