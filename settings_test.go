package settings

import (
	"testing"
)

func TestGetOsVariable(t *testing.T) {
	res := GetOsVariable("hello", "hello")
	if "hello" != GetOsVariable("hello", "hello") {
		t.Error("Expected hello, got ", res)
	}
}

func TestReadConfigFile(t *testing.T) {
	type Config struct {
		Name string `json:"name"`
	}

	var config Config

	_, errNoFile := ReadConfigFile("./test123.json", &config)

	if errNoFile == nil {
		t.Error("Should fail as no config file exists")
	}

	_, err := ReadConfigFile("./test.json", &config)

	if err != nil {
		t.Error("Where is Config file???")
	}

	if config.Name != "test" {
		t.Error("Parse error, expected test, got ", config.Name)
	}
}
