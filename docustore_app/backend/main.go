package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pathogende/docustore/pkg/api"
	"github.com/pathogende/docustore/pkg/config"
	"github.com/pathogende/docustore/pkg/database"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func main() {
	err := config.Init()
	if err != nil {
		logger.Fatalf("Error initializing config: %v", err)
	}
	logger.Debug("Initialized config")
	mongoClient, err := database.NewDB()
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}

	// Set up API
	r := api.Router(mongoClient)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start API server in a goroutine
	go func() {
		logger.Info("Serving api..")
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Error starting API server: %v", err)
		}

	}()

	// Wait for an interrupt signal from the OS
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Shut down the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Error shutting down API server: %v", err)
	}
}

func init() {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
}
