package main

import (
	"context"
	"dota-api/internal"
	db2 "dota-api/internal/db"
	"dota-api/internal/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	logger := internal.InitializeLogger()
	db := db2.NewDB(logger)
	mux := routes.InitializeRoutes(logger, db)
	server := &http.Server{
		Addr:     ":4000",
		ErrorLog: logger.ErrorLog,
		Handler:  mux,
	}

	errs := make(chan error, 2)

	go func() {
		logger.InfoLog.Println("RUNNING API")
		err := server.ListenAndServe()
		if err != nil {
			errs <- err
			return
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("Got the signal", <-c)
	}()

	logger.ErrorLog.Print("terminating api with error", <-errs)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
