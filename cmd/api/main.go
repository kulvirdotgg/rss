package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	env "github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// read the ENV vars from .env file
	env.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("No/Incorrect PORT provided in the env vars")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	cfg := config{
		port: port,
		env:  env,
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
