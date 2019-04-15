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
	// TODO continue here, I need to kill gin whenever I want
	//srv := &http.Server{
	//	Addr:    ":" + os.Getenv("SERVICE_PORT"),
	//	Handler: engine,
	//}
	//go func() {
	//	// service connections
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Printf("listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Println("Shutdown Server ...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}

}
