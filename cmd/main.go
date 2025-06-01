package main

import (
	"log"
	"strings"

	"github.com/Anvinalias/az-blob-downloader/internal/config"
	"github.com/Anvinalias/az-blob-downloader/internal/decrypt"
	"github.com/Anvinalias/az-blob-downloader/internal/request"
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

	requests, err := request.ReadRequests("request.txt")
	if err != nil {
		log.Fatalf("Failed to read requests: %v", err)
	}
	for _, req := range requests {
		blobs, err := storage.ListBlobsWithPrefix(client, cfg.Storage.BlobName, req.Prefix)
		if err != nil {
			log.Printf("Failed to list blobs for prefix %s: %v", req.Prefix, err)
			continue
		}
		// To print the path steps:
		baseNames, err := storage.BuildShortestUpgradePath(blobs, req)
		if err != nil {
			log.Printf("%v", err)
			continue
		}
		if len(baseNames) == 1 {
			log.Printf("Shortest path for %s: [%s]", req.Raw, baseNames[0])
		} else {
			log.Printf("Shortest path for %s: [%s]", req.Raw, strings.Join(baseNames, " -> "))
		}
		err = storage.DownloadBlobsByStep(client, cfg.Storage.BlobName, blobs, baseNames, cfg.Paths.DownloadPath)
		if err != nil {
			log.Fatalf("Download failed: %v", err)
		}
	}
	return nil
}
