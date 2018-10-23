package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
	"github.com/pedromsmoreira/go-simple-rest-api/database"
	"github.com/pedromsmoreira/go-simple-rest-api/handlers"
)

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	loader := configurations.JSONLoader{Fs: configurations.OsFS{}}
	config, err := loader.Load()

	redisDb := &database.RedisRepository{
		Client: database.CreateClient(config.Redis),
	}

	if err != nil {
		log.Panic("Error occurred loading configs.")
		panic(err)
	}

	hcHandler := handlers.NewHealthCheckHandler(redisDb)

	router := handlers.CreateRoutes(hcHandler)

	srv := &http.Server{
		Addr:         config.App.Address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Listen and serve without blocking
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.

	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
