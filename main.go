package main

import (
	"github.com/insmnia/go-users-service/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app, err := cmd.InitApp()
	if err != nil {
		log.Fatalf("Cannot init app due to %s", err.Error())
	}

	log.Print("Server started")
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
