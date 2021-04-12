package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func UploadEnvironmentVariables(pathToConfigFile string) {
	ymlFile, err := os.Open(pathToConfigFile)
	if err != nil {
		panic(err)
	}

	defer ymlFile.Close()

	var variables = make(map[string]string)

	byteValue, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(byteValue, &variables)
	if err != nil {
		panic(err)
	}

	for k, v := range variables {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}
