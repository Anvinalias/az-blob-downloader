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

// DownloadMatchingBlobs downloads all blobs from the specified container whose names contain the given pattern.
// Temporary test function
func DownloadMatchingBlobs(client *azblob.Client, containerName, pattern, downloadDir string) error {
	ctx := context.Background()
	pager := client.NewListBlobsFlatPager(containerName, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to list blobs: %w", err)
		}

		for _, blob := range page.Segment.BlobItems {
			if strings.Contains(*blob.Name, pattern) {
				log.Printf("Downloading blob: %s", *blob.Name)
				err := downloadBlob(ctx, client, containerName, *blob.Name, downloadDir)
				if err != nil {
					return fmt.Errorf("failed to download blob %s: %w", *blob.Name, err)
				}
			}
		}
	}

	return nil
}

func downloadBlob(ctx context.Context, client *azblob.Client, containerName, blobName, downloadDir string) error {
	resp, err := client.DownloadStream(ctx, containerName, blobName, nil)
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
