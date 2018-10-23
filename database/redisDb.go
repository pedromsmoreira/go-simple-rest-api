package database

import (
	"github.com/go-redis/redis"
	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
)

type RedisRepository struct {
	Client *redis.Client
}

func CreateClient(c configurations.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.ConnectionString,
		Password: c.Password,
		DB:       0,
	})
}

/*
func (r RedisRepository) Get(key string) (string, error) {

}

func (r RedisRepository) Set(key, value string, expiration time.Duration) error {

}
*/

func (r RedisRepository) Ping() (string, error) {
	pong, err := r.Client.Ping().Result()

	if err != nil {
		return "", err
	}

	return pong, nil
}
