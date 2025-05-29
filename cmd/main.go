package main

import (
	"log"

	"github.com/Anvinalias/az-blob-downloader/internal/config"
	"github.com/Anvinalias/az-blob-downloader/internal/decrypt"
	"github.com/Anvinalias/az-blob-downloader/internal/storage"
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

	client, err := storage.NewClient(connStr)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	log.Println("Azure Blob client created successfully")

	// err = storage.DownloadMatchingBlobs(client, cfg.Storage.BlobName, "lpsqpalapp-1.0.0.0-4.0.0.0", cfg.Paths.DownloadPath)
	// if err != nil {
	// 	log.Fatalf("Download failed: %v", err)
	// }

	// test
	blobs, err := storage.ListMatchingBlobs(client, cfg.Storage.BlobName, "maintenancepalapp")
	if err != nil {
		log.Fatal(err)
	}
	for _, blob := range blobs {
		log.Println("Matched:", blob)
	}

	return nil
}
