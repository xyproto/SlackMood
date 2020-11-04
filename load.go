package emojimood

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Config config

func LoadConfig(filePath string) error {
	Config = config{}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config file not found: " + filePath)
	}

	configContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configContent, &Config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %s", err)
	}

	return nil
}
