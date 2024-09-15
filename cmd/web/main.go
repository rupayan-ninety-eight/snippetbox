package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rupayan-ninety-eight/snippetbox/internal/models"
)

var (
	port string
)

type config struct {
	addr      int
	staticDir string
}

type application struct {
	config        config
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	cfg := config{
		addr:      4000,
		staticDir: "./ui/static",
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, database)

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

	db, err := openDB(dsn)
	if err != nil {
		logger.Error((err.Error()))
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config:        cfg,
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("starting server", slog.Int("addr", cfg.addr))
	err = http.ListenAndServe(":"+strconv.Itoa(cfg.addr), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
