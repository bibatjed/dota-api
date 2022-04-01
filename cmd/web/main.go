package main

import (
	"context"
	"dota-api/internal"
	"dota-api/internal/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger := internal.InitializeLogger()
	mux := routes.InitializeRoutes(logger)
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
