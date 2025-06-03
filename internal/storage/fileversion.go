package storage

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type UploadInfo struct {
	AppName     string `json:"appName"`
	Version     string `json:"version"`
	Criticality string `json:"criticality"`
	CheckSum    any    `json:"checkSum"`
}

// GenerateUploadedVersion scans a step directory and creates uploadedversion.txt
func GenerateUploadedVersion(stepDir string) error {
	entries, err := os.ReadDir(stepDir)
	if err != nil {
		return fmt.Errorf("failed to read step directory: %w", err)
	}

	seen := make(map[string]struct{})
	var result []UploadInfo

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		base := strings.TrimSuffix(name, filepath.Ext(name))
		if strings.HasSuffix(base, "-release") {
			continue // skip sidecar files
		}

		if _, exists := seen[base]; exists {
			continue
		}
		seen[base] = struct{}{}

		parts := strings.Split(base, "-")
		if len(parts) < 3 {
			continue
		}

		appName := strings.Join(parts[:len(parts)-2], "-")
		version := parts[len(parts)-1]

		result = append(result, UploadInfo{
			AppName:     appName,
			Version:     version,
			Criticality: "",
			CheckSum:    nil,
		})
	}

	jsonPath := filepath.Join(stepDir, "uploadedversion.txt")
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to generate JSON for uploadedversion.txt: %w", err)
	}
	if err := os.WriteFile(jsonPath, data, fs.FileMode(0644)); err != nil {
		return fmt.Errorf("failed to write uploadedversion.txt: %w", err)
	}
	var appNames []string
	for _, info := range result {
		appNames = append(appNames, info.AppName)
	}
	fmt.Printf("uploadedversion.txt updated with applications: %s\n", strings.Join(appNames, ", "))

	return nil
}
