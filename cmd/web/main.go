package main

import (
	"log"
	"net/http"
	"strconv"
)

var (
	port string
)

type config struct {
	addr      int
	staticDir string
}

func main() {
	cfg := config{
		addr:      4000,
		staticDir: "./ui/static",
	}

	if port != "" {
		addr, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("invalid port %v", err)
		}
		cfg.addr = addr
	}

	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %d", cfg.addr)
	err := http.ListenAndServe(":"+strconv.Itoa(cfg.addr), mux)
	log.Fatal(err)
}
