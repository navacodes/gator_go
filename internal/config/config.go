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

// SetUser updates the current username in the config and saves it back to the file

func (c *Config) SetUser(username string) error {
	// update the username field memory
	c.CurrentUsername = username

	// get home directory to write back to the file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// build the config path
	configPath := filepath.Join(homeDir, ".gatorconfig.json")

	//marshal(returns json as a byte slice) the updated config to json
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	//write the updated json data back to the file
	// 0600 is the file permission(read and write for the owner only)
	err = os.WriteFile(configPath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
