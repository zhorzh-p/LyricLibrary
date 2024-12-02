package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/zhorzh-p/LyricLibrary/docs"
	"github.com/zhorzh-p/LyricLibrary/internal/infrastructure/server"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Fatalf("Error loading .env file")
	}

	router, err := server.NewServer()
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to initialize server")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		logrus.Warn("Environment variable PORT not set. Using default port 8000.")
	}

	logrus.WithField("port", port).Info("Starting server")
	if err := router.Run(":" + port); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}
