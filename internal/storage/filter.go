package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// ListMatchingBlobs lists blob names in a container that contain the specified pattern.
// Temporary test function
func ListMatchingBlobs(client *azblob.Client, containerName, pattern string) ([]string, error) {
	ctx := context.Background()
	pager := client.NewListBlobsFlatPager(containerName, nil)

	var matches []string

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list blobs: %w", err)
		}

		for _, blob := range page.Segment.BlobItems {
			if strings.Contains(*blob.Name, pattern) {
				matches = append(matches, *blob.Name)
			}
		}
	}

	return matches, nil
}
