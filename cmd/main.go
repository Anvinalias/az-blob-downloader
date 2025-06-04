package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Anvinalias/az-blob-downloader/internal/config"
	"github.com/Anvinalias/az-blob-downloader/internal/decrypt"
	"github.com/Anvinalias/az-blob-downloader/internal/logging"
	"github.com/Anvinalias/az-blob-downloader/internal/request"
	"github.com/Anvinalias/az-blob-downloader/internal/storage"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("ERROR: could not determine executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	configPath := findFile("config.yaml", exeDir)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("ERROR: Failed to load config: %v", err)
	}
	logFile, err := logging.Setup(cfg.Paths.LogPath)
	if err != nil {
		log.Fatalf("ERROR: Failed to set up logging: %v", err)
	}
	defer logFile.Close()

	if err := run(cfg, exeDir); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}

func run(cfg *config.Config, exeDir string) error {

	// Decrypt the encrypted connection string
	connStr, err := decrypt.DecryptAESGCM(cfg.Storage.ConnectionStringEncrypted, cfg.Storage.Passphrase)
	if err != nil {
		return wrapErr("decrypting connection string", err)
	}

	client, err := storage.NewClient(connStr)
	if err != nil {
		return wrapErr("creating Azure Blob client", err)
	}
	log.Println("Azure Blob client created successfully")

	requestPath := findFile("request.txt", exeDir)
	requests, err := request.ReadRequests(requestPath)
	if err != nil {
		return wrapErr("reading requests", err)
	}
	for _, req := range requests {
		blobs, err := storage.ListBlobsWithPrefix(client, cfg.Storage.BlobName, req.Prefix)
		if err != nil {
			log.Printf("ERROR: Failed to list blobs for prefix %s: %v", req.Prefix, err)
			continue
		}
		// To print the path steps:
		baseNames, err := storage.BuildShortestUpgradePath(blobs, req)
		if err != nil {
			log.Printf("ERROR: Failed to build path for %s: %v", req.Raw, err)
			continue
		}
		if len(baseNames) == 1 {
			log.Printf("Shortest path for %s: [%s]", req.Raw, baseNames[0])
		} else {
			log.Printf("Shortest path for %s: [%s]", req.Raw, strings.Join(baseNames, " -> "))
		}
		err = storage.DownloadBlobsByStep(client, cfg.Storage.BlobName, blobs, baseNames, cfg.Paths.DownloadPath)
		if err != nil {
			log.Printf("ERROR: Download failed for %s: %v", req.Raw, err)
			continue
		}
		log.Printf("Downloaded %s", req.Raw)
	}
	return nil
}

func wrapErr(context string, err error) error {
	return fmt.Errorf("%s: %w", context, err)
}

func findFile(filename string, exeDir string) string {
	cwd, _ := os.Getwd()
	cwdPath := filepath.Join(cwd, filename)
	if _, err := os.Stat(cwdPath); err == nil {
		return cwdPath
	}
	exePath := filepath.Join(exeDir, filename)
	if _, err := os.Stat(exePath); err == nil {
		return exePath
	}
	return cwdPath // Will error later if not found
}
