package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ParcelsApiToken string `json:"parcelsApiToken"`
}

func ReadConfig(filePath string) (*Config, error) {
	// Get the absolute path to the file
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	// Read the file
	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
