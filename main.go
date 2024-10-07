package main

import (
	"os"
	"log/slog"
	"github.com/go-chi/chi/v5"
	"net/http"
	"fmt"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME ROUT")
}

func main() {
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
        slog.Error("Unable to open a file for writing")
    }

	opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

	logger := slog.New(slog.NewJSONHandler(f, opts))
	logger.Info("Info message")

	// ConnectDatabase(logger)

	router := chi.NewRouter()
    router.HandleFunc("/", handlerHome)

	err = http.ListenAndServe(os.Getenv("PORT"), router)
	if err != nil {
		logger.Error("failed start server")
		panic(err)
	}
}