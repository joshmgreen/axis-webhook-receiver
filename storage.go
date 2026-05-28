package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func EnsureOutputDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

func SaveImage(outputDir string, imageData []byte, filename string) (string, error) {
	if err := EnsureOutputDir(outputDir); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Generate timestamped filename if not provided
	if filename == "" {
		filename = fmt.Sprintf("image_%s.jpg", time.Now().Format("2006-01-02T15-04-05.000Z07-00"))
	} else {
		// Preserve original filename with timestamp prefix
		ext := filepath.Ext(filename)
		name := filename[:len(filename)-len(ext)]
		filename = fmt.Sprintf("%s_%s%s", name, time.Now().Format("2006-01-02T15-04-05.000Z07-00"), ext)
	}

	fullPath := filepath.Join(outputDir, filename)
	err := os.WriteFile(fullPath, imageData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write image: %w", err)
	}

	return fullPath, nil
}
