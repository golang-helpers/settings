package settings

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

const (
	FILE_NOT_FOUND = "file not found"
)

func GetOsVariable(name string, defaultValue string) string {
	variable := os.Getenv(name)

	if variable == "" {
		return defaultValue
	}

	return variable
}

func ReadConfigFile(fileName string, configInterface interface{}) (interface{}, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("File error: %v\n", err)
		return nil, errors.New(FILE_NOT_FOUND)
	}
	json.Unmarshal(file, &configInterface)

	return configInterface, nil
}
