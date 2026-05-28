package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request, cfg Config) {
	log.Println("\n=== INCOMING WEBHOOK ===")
	log.Printf("Method: %s\n", r.Method)
	log.Printf("URL: %s\n", r.RequestURI)
	log.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))

	// Log headers
	log.Println("Headers:")
	for key, values := range r.Header {
		for _, val := range values {
			log.Printf("  %s: %s\n", key, val)
		}
	}

	// Parse multipart form (32 MB max)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		log.Printf("Multipart parse error: %v (might be regular form data)\n", err)
	}

	// Log form fields
	if r.MultipartForm != nil && r.MultipartForm.Value != nil {
		log.Println("Form Fields:")
		for key, values := range r.MultipartForm.Value {
			for _, val := range values {
				log.Printf("  %s: %s\n", key, val)
			}
		}
	}

	// Extract and save files (images + event data)
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		log.Println("Files:")
		for fieldName, files := range r.MultipartForm.File {
			log.Printf("  Field: %s\n", fieldName)
			for _, fileHeader := range files {
				log.Printf("    Filename: %s (size: %d bytes)\n", fileHeader.Filename, fileHeader.Size)

				// Open file from multipart
				file, err := fileHeader.Open()
				if err != nil {
					log.Printf("    Error opening file: %v\n", err)
					continue
				}
				defer file.Close()

				// Read file data
				fileData, err := io.ReadAll(file)
				if err != nil {
					log.Printf("    Error reading file: %v\n", err)
					continue
				}

				// Log file contents (for JSON events)
				if strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".json") {
					log.Printf("    JSON Content:\n%s\n", string(fileData))
				}

				// Check if this is an image
				isImage := isImageFile(fileHeader.Filename)
				if isImage {
					// Save image to disk
					savedPath, err := SaveImage(cfg.OutputDir, fileData, fileHeader.Filename)
					if err != nil {
						log.Printf("    Error saving image: %v\n", err)
					} else {
						log.Printf("    ✓ Saved to: %s\n", savedPath)
					}
				} else if strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".json") {
					log.Printf("    ✓ Event JSON logged above\n")
				}
			}
		}
	}

	log.Println("=== END WEBHOOK ===\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received and logged"))
}

func isImageFile(filename string) bool {
	lowerName := strings.ToLower(filename)
	imageExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	for _, ext := range imageExts {
		if strings.HasSuffix(lowerName, ext) {
			return true
		}
	}
	return false
}
