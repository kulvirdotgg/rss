package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rss/internal/logger"
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
	logger *logger.Logger
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

	logger := logger.New(os.Stdout, logger.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err = srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}
