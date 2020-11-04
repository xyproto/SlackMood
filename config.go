package emojimood

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the structure in the config.yml file
type Config struct {
	SlackToken string `yaml:"slack_token"`
	DBPath     string `yaml:"db_path"`
	RankFile   string `yaml:"rank_file"`
}

// LoadConfig loads a config.yml file into a Config struct
func LoadConfig(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: " + filePath)
	}

	configContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %s", err)
	}

	return &config, nil
}
