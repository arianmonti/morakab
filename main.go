package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"time"

	"morakab/config"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	conn, _ := sql.Open("postgres", config.Cfg.DatabaseURL)
	if err := conn.Ping(); err != nil {
		panic(err)
	}
	defer conn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	file, _ := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	multilogged := io.MultiWriter(file, os.Stdout)
	logged := handlers.LoggingHandler(multilogged, mux)
	server := &http.Server{
		Addr:         ":5000",
		Handler:      logged,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
