package main 

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        slog.Error(fmt.Sprintf("Error loading .env file: %v", err))
    } 

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Liberty Hills app!"))
	})

	// serve traffic
	port := os.Getenv("PORT")
	slog.Info(fmt.Sprintf("Server starting on localhost:%s...", port))
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}