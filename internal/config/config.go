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
	// downloadPath is required
	if cfg.Paths.DownloadPath == "" {
		return nil, errors.New("downloadPath is required in config.yaml")
	}
	if err := ensureDir(cfg.Paths.DownloadPath); err != nil {
		return nil, err
	}

	// Always use default log path
	cfg.Paths.LogPath = filepath.Join(".", "logs")
	if err := ensureDir(cfg.Paths.LogPath); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// makes sure the directory exists
func ensureDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}
