package main 

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received a request to '/'")
		w.Write([]byte("Welcome to the Liberty Hills app!"))
	})

	// serve traffic
	slog.Info("Server starting on localhost:3001 ...")
	http.ListenAndServe(":3001", r)
}