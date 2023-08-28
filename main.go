package main

import (
	"os"
	"net/http"
	"log/slog"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	logger := setLogLevel()
	check(err)

	files, err := os.ReadDir("client")
	check(err)

	for _, e := range files {
		slog.Debug("Client File: %s", e.Name())
	}

	http.HandleFunc("/", root)

	slog.Info("Starting Server on Port 3000...")
	http.ListenAndServe(":3000", nil)
}

func setLogLevel() *slog.Logger {
	lvl, lvl_found := os.LookupEnv("LOG_LEVEL")
	logger := new(slog.LevelVar)

	if lvl_found {
		switch lvl {
			case "DEBUG": logger.Set(slog.LevelDebug)
			case "WARN": logger.Set(slog.LevelWarn)
			
			case "ERROR": logger.Set(slog.LevelError)
		}
	}

	return logger
}
// --- HANDLERS ---

func root(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Route Hit:  ROOT")

	http.ServeFile(w, r, "client/index.html")
}
