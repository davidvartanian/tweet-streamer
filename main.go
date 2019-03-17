package main

import (
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	_ "main/docs"
	"main/server"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	engine := server.GetServer()
	err := engine.Run(":" + os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Errorf("Error trying to run server: %v", err)
	}
}
