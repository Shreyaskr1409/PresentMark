package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shreyaskr1409/PresentMark/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	l := log.New(os.Stdout, "server: ", log.LstdFlags)
	l.Println("Hello world")

	defaultRouter := mux.NewRouter()
	defaultRouter.UseEncodedPath()
	defaultRouter.Use(mux.CORSMethodMiddleware(defaultRouter))
	defaultRouter.Use(middlewares.LoggingMiddleware(l))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      defaultRouter,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Println(err)
		}
	}()
	l.Println("Server listening at port: ", 9090)

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdownChannel
	l.Println("Recieved a shutdown signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	l.Println("Shutting down server...")
	s.Shutdown(ctx)
	l.Println("Shutdown successful")
}
