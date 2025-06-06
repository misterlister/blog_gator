package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()

	if err != nil {
		return nil, fmt.Errorf("could not locate config file: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config JSON: %w", err)
	}

	return &cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("could not locate config file: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return fmt.Errorf("could not marshal config to JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("could not write to config file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}
