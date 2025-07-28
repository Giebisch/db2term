package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// Points to UserConfigDir/db2term/config.json
// Creates a new file if it does not exist yet
func GetConfigFile() (*os.File, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		slog.Error("Can't get UserConfigDir", "error", err)
		return nil, err
	}

	err = EnsureDirExists(filepath.Join(configDir, "db2term"))
	if err != nil {
		slog.Error("Error creating directory", "error", err)
		return nil, err
	}

	file_path := filepath.Join(configDir, "db2term", "config.json")

	file, err := os.OpenFile(
		file_path, os.O_RDWR|os.O_CREATE, 0644)
	if errors.Is(err, os.ErrNotExist) {
		// file, err = os.CreateTemp(os.TempDir(), "db2term_config.json")
		if err != nil {
			return file, fmt.Errorf("unable to create config file: %w", err)
		}
	}
	return file, nil
}

// Reads the config file
func ReadConfig(file *os.File) (Config, error) {
	var config Config

	bytes, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Config ReadAll Error", "error", err)
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		slog.Error("Config Unmarshal Error", "error", err)
		return config, err
	}
	defer file.Close()

	return config, nil
}

// Saves the current config to `db2term_config.json`
func (c *Config) WriteConfig(file *os.File) error {
	return nil
}
