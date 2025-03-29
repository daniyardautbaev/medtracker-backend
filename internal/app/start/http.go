package start

import (
	"context"
	"fmt"
	"log"
	"medtracker/medtracker/internal/app/config"
	"net/http"
)

func HTTP(ctx context.Context, cfg *config.Config) {
	// Пример создания HTTP сервера
	mux := http.NewServeMux()

	// Пример health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%s", cfg.HTTPServer.Port)
	log.Printf("HTTP server running on %s", addr)

	// Запуск HTTP сервера
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
