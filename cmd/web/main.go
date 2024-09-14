package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var (
	port string
)

type config struct {
	addr      int
	staticDir string
}

type application struct {
	logger *slog.Logger
	config config
}

func main() {
	cfg := config{
		addr:      4000,
		staticDir: "./ui/static",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	if port != "" {
		addr, err := strconv.Atoi(port)
		if err != nil {
			logger.Error("invalid port", "error", err.Error())
		}
		if err == nil {
			cfg.addr = addr
		}
	}

	app := &application{
		logger: logger,
		config: cfg,
	}

	logger.Info("starting server", slog.Int("addr", cfg.addr))
	err := http.ListenAndServe(":"+strconv.Itoa(cfg.addr), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
