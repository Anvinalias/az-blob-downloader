package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// NewClient creates a new Azure Blob Storage client using the given connection string.
func NewClient(connectionString string) (*azblob.Client, error) {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
