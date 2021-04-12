package main

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/SKilliu/novels-backend/internal/server"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const pathToConfigFile = "./envs.yaml"

// @title App API
// @version 1.0
// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	log := logrus.New()
	logger := logrus.NewEntry(log)

	utils.UploadEnvironmentVariables(pathToConfigFile)

	db.Init(logger)
	//s3.Init(logger)
	server.Init(logger)

	err := server.Start()
	if err != nil {
		panic(errors.Wrap(err, "failed to start api"))
	}
}
