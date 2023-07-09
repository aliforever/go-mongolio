package mongolio

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URI      string `json:"uri"`
	Database string `json:"database"`
}

func LoadFromEnv() (*Config, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}

	database := os.Getenv("MONGO_DATABASE")
	if database == "" {
		return nil, fmt.Errorf("MONGO_DATABASE is not set")
	}

	return &Config{
		URI:      uri,
		Database: database,
	}, nil
}

// LoadFromJSONFile loads the config from a JSON file.
func LoadFromJSONFile(path string) (*Config, error) {
	var config *Config

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// LoadFromJSONString loads the config from a JSON string.
func LoadFromJSONString(jsonString string) (*Config, error) {
	var config *Config

	err := json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
