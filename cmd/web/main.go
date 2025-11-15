package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/StevanFreeborn/links.stevanfreeborn.com/internal/handlers"
	"github.com/StevanFreeborn/links.stevanfreeborn.com/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/css/", handlers.CSS)
	mux.HandleFunc("/fonts/", handlers.Fonts)

	loggingMux := middleware.Logging(mux)

	const addr = "localhost:7777"

	server := &http.Server{
		Addr:    addr,
		Handler: loggingMux,
	}

	go func() {
		log.Printf("Server listening at http://%s\n", addr)

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}

		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)

	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	log.Println("Graceful shutdown complete.")
}
