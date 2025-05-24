package main

import (
	"log"

	"github.com/Anvinalias/az-blob-downloader/config"
	"github.com/Anvinalias/az-blob-downloader/internal/decrypt"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return err
	}

	// Decrypt the encrypted connection string
	connStr, err := decrypt.DecryptAESGCM(cfg.Storage.ConnectionStringEncrypted, cfg.Storage.Passphrase)
	if err != nil {
		return err
	}

	log.Printf("Blobname: %s", cfg.Storage.BlobName)
	log.Printf("Download path: %s", cfg.Paths.DownloadPath)
	log.Printf("Log path: %s", cfg.Paths.LogPath)

	return nil
}
