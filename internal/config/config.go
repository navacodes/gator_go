package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

//Read functions

func Read() (Config, error) {
	//Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	configPath := filepath.Join(homeDir, ".gatorconfig.json")
	//read file contents

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	//unmarshal json data
	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
