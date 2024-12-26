// main.go
// Greenlight movie database app from the Let's Go Further! ebook

package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Declare a string containing the app version number.
const version = "1.0.0"

// Define a config struct to hold the configuration settings for our app.
type config struct {
	port int
	env  string
}

/* Define an application struct to hold the dependencies
   for our HTTP handlers, helpers, and middleware. */
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new structured logger which writes log entries to stdout.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the application struct.
	app := &application{
		config: cfg,
		logger: logger,
	}

    // Use the httprouter instance returned by app.routes() as the server handler.
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      app.routes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }

	// Start the HTTP server.
	logger.Info("Starting HTTP server on port:", "addr", srv.Addr, "env", cfg.env)
    err := srv.ListenAndServe()
    logger.Error(err.Error())
    os.Exit(1)
}
