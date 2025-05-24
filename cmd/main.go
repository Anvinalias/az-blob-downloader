package main

import (
	"log"

	"github.com/Anvinalias/az-blob-downloader/config"
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
	log.Printf("Blobname: %s", cfg.Storage.BlobName)
	log.Printf("Download path: %s", cfg.Paths.DownloadPath)
	log.Printf("Log path: %s", cfg.Paths.LogPath)

	return nil
}
