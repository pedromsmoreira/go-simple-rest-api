package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedromsmoreira/go-simple-rest-api/database"
	"github.com/pedromsmoreira/go-simple-rest-api/model"
)

type HealthCheckHandler struct {
	RedisDb database.Repository
	Routes  Routes
}

func NewHealthCheckHandler(redis database.Repository) *HealthCheckHandler {
	return &HealthCheckHandler{
		RedisDb: redis,
	}
}

func (hc *HealthCheckHandler) Shallow(w http.ResponseWriter, r *http.Request) {
	hcs := []model.Shallow{redisPing(hc.RedisDb)}

	resp, err := json.Marshal(&hcs)
	if err != nil {
		http.Error(w, "Error Converting to json.", http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(resp))
}

func (hc *HealthCheckHandler) Deep(w http.ResponseWriter, r *http.Request) {
	// redis -> number of keys, timetaken
	// mongo -> number of documents in database, timetaken
}

func redisPing(redis database.Repository) model.Shallow {
	pong, err := redis.Ping()

	if err != nil {
		return model.NewShallow("Redis", "Server is down.", false)
	}

	return model.NewShallow("Redis", pong, true)
}

func CreateHandlerRoutes(hc *HealthCheckHandler) Routes {
	return Routes{
		Route{
			Method:      "GET",
			Name:        "Shallow Healthchecks",
			Pattern:     "/healthchecks/shallow",
			HandlerFunc: hc.Shallow,
		},
	}
}
