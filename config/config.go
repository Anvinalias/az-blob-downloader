package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Structure of config.yaml
type Config struct {
	Storage StorageConfig `yaml:"storage"`
	Paths   PathConfig    `yaml:"paths"`
}

type StorageConfig struct {
	ConnectionStringEncrypted string `yaml:"connectionStringEncrypted"`
	Passphrase                string `yaml:"passphrase"`
	BlobName                  string `yaml:"blobName"`
}

type PathConfig struct {
	DownloadPath string `yaml:"downloadPath"`
	LogPath      string `yaml:"-"`
}

// reads and parses the config file
func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Validate storage config fields
	if cfg.Storage.ConnectionStringEncrypted == "" {
		return nil, errors.New("storage.connectionStringEncrypted is required in config.yaml")
	}
	if cfg.Storage.Passphrase == "" {
		return nil, errors.New("storage.passphrase is required in config.yaml")
	}
	if cfg.Storage.BlobName == "" {
		return nil, errors.New("storage.blobName is required in config.yaml")
	}

	// Always use default log path
	cfg.Paths.LogPath = filepath.Join(".", "logs")

	return &cfg, nil
}
