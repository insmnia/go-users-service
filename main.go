package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := gin.Default()
	log.Print("Service started")
	go func() {
		err := app.Run()
		if err != nil {
			log.Fatalf("Failed to start server due to %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("Shutting down server")
}