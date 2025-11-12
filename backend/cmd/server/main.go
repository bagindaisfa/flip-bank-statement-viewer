package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bagindaisfa/flip-bank-statement-viewer/internal/handler"
	"github.com/bagindaisfa/flip-bank-statement-viewer/internal/repository"
	"github.com/bagindaisfa/flip-bank-statement-viewer/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/upload", h.Upload)
	mux.HandleFunc("/balance", h.Balance)
	mux.HandleFunc("/issues", h.Issues)

	handlerWithCORS := enableCORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handlerWithCORS); err != nil {
		log.Fatal(err)
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Untuk permintaan OPTIONS (preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
