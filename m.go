package main

import (
	"link-ly/api"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		Addr:        ":8080",
		Handler:     handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
