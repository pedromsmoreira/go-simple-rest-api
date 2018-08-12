package main

import (
	"log"
	s "strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
	"github.com/pedromsmoreira/go-simple-rest-api/database"
	"github.com/pedromsmoreira/go-simple-rest-api/healthcheck"
)

func main() {

	loader := configurations.JSONLoader{Fs: configurations.OsFS{}}
	config, err := loader.Load()

	redisDb := &database.RedisRepository{
		Client: database.CreateClient(config.Redis),
	}

	hcController := healthcheck.NewHealthCheckController(redisDb)

	if err != nil {
		log.Panic("Error occurred loading configs.")
		panic(err)
	}

	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	healthCheckAPI := app.Party("/healthchecks")
	{
		healthCheckAPI.Get("/shallow", hcController.Shallow)

		healthCheckAPI.Get("/deep", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"Message": "Deep Healthcheck"})
		})
	}

	app.Run(iris.Addr(s.Join([]string{":", config.App.Port}, "")), iris.WithoutServerError(iris.ErrServerClosed))
}
