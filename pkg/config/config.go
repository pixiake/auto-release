package config

import (
	"github.com/pixiake/auto-release/pkg/repo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func GetConfig() (*repo.ChangeLog, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configPath := currentDir + "/config.yaml"
	config, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	c := repo.ChangeLog{}

	err = yaml.Unmarshal(config, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
