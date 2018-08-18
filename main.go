package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
	"github.com/pedromsmoreira/go-simple-rest-api/database"
	"github.com/pedromsmoreira/go-simple-rest-api/handlers"
)

func main() {

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

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), router); err != nil {
		log.Fatal(err)
	}
}
