package main

import (
	"os"
)

type Config struct {
	Port      string
	OutputDir string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "./images"
	}

	return Config{
		Port:      port,
		OutputDir: outputDir,
	}
}
