package main

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/SKilliu/novels-backend/internal/server"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const pathToConfigFile = "./envs.yaml"

// @title Novels REST API
// @version 0.0.2
// @description REST API for Novels app.
// @description New in version:<br> - socials sign up was deleted. Now we have 1 endpoint for signin/signup.<br> - some minor fixes
// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
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
