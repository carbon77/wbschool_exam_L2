package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ru/zakat/server/router"
	"syscall"
	"time"
)

func main() {
	host := "127.0.0.1"
	port, ok := os.LookupEnv("DEV12_PORT")
	if !ok {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)
	server := &http.Server{
		Addr: address,
	}

	router.InitRouter()

	go func() {
		log.Printf("Server is listening on %s port\n", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v\n", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v\n", err)
	}
	log.Println("Server stopped")
}
