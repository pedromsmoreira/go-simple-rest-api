package testing

import (
	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
)

var MockDevConfig = configurations.Configuration{
	App: configurations.AppConfig{Port: "1"},
	Redis: configurations.RedisConfig{
		ConnectionString: "testconn",
		Password:         "testPassword",
		User:             "testUser",
		Port:             1,
	},
	MongoDb: configurations.MongoConfig{
		ConnectionString: "testconn",
		Password:         "testPassword",
		User:             "testUser",
		Port:             1,
		DbName:           "testDbName",
	},
}
