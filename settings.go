package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetOsVariable(name string, defaultValue string) string {
	variable := os.Getenv(name)

	if variable == "" {
		return defaultValue
	}

	return variable
}

func ReadConfigFile(fileName string, configInterface interface{}) interface{} {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &configInterface)

	return configInterface
}
