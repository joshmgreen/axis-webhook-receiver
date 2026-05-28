package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := LoadConfig()

	log.Printf("Starting Axis webhook receiver on port %s\n", cfg.Port)
	log.Printf("Output directory: %s\n", cfg.OutputDir)

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		WebhookHandler(w, r, cfg)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Listening on http://0.0.0.0:%s\n", cfg.Port)
	log.Println("POST /webhook - Receive Axis camera events")
	log.Println("GET  /health  - Health check")

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
