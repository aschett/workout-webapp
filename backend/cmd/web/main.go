package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	addr := flag.String("addr", ":4000", "Http network address")

	logger.Info("Starting Server on", "addr", *addr)

	app := &application{
		logger: logger,
	}

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
}
