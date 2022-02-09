package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"time"

	"morakab/config"
	mhand "morakab/handlers"
	"morakab/pkg"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	conn, _ := sql.Open("postgres", config.Cfg.DatabaseURL)
	if err := conn.Ping(); err != nil {
		panic(err)
	}
	defer conn.Close()
	DB = conn

	mux := http.NewServeMux()

	morakab := pkg.Morakab{DB: DB}
	handler := &mhand.HTTPHandler{Morakab: &morakab}
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)

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
