package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// DownloadBlobsByStep downloads all blobs whose names start with any baseName.
// Each upgrade step may have multiple related blobs with different extensions (e.g., .zip, .z01, .z02)
// or suffixes (e.g., -release.txt).
func DownloadBlobsByStep(client *azblob.Client, containerName string, allBlobs []string, baseNames []string, downloadPath string) error {
	for i, base := range baseNames {
		stepDir := filepath.Join(downloadPath, containerName, fmt.Sprintf("step%d", i+1))
		if err := os.MkdirAll(stepDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", stepDir, err)
		}
		for _, blob := range allBlobs {
			if strings.HasPrefix(blob, base) {
				log.Printf("Downloading: %s", blob)
				if err := downloadBlob(client, containerName, blob, stepDir); err != nil {
					return fmt.Errorf("failed to download %s: %w", blob, err)
				}
			}
		}
		// Generate uploadedversion.txt for this step
		if err := GenerateUploadedVersion(stepDir); err != nil {
			log.Printf("Failed to generate uploadedversion.txt for %s: %v", stepDir, err)
		}
	}
	return nil
}

// downloadBlob downloads a single blob and saves it locally.
func downloadBlob(client *azblob.Client, containerName, blobName, downloadDir string) error {
	resp, err := client.DownloadStream(context.Background(), containerName, blobName, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Save the blob locally
	filePath := filepath.Join(downloadDir, filepath.Base(blobName))
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	return err
}
