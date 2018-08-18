package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedromsmoreira/go-simple-rest-api/database"
	"github.com/pedromsmoreira/go-simple-rest-api/model"
)

type HealthCheckHandler struct {
	RedisDb database.Repository
}

func NewHealthCheckHandler(redis database.Repository) *HealthCheckHandler {
	return &HealthCheckHandler{
		RedisDb: redis,
	}
}

func (hc *HealthCheckHandler) Shallow(w http.ResponseWriter, r *http.Request) {
	hcs := make([]model.Shallow, 0)
	hcs = append(hcs, redisPing(hc.RedisDb))

	if err := json.NewEncoder(w).Encode(hcs); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
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
