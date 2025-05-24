package config

import (
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

	// Always use default log path
	cfg.Paths.LogPath = filepath.Join(".", "logs")

	return &cfg, nil
}
