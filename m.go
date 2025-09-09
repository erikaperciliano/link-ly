package main

import (
	"errors"
	"log/slog"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	return errors.New("foo")
}
