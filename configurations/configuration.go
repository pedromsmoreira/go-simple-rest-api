package configurations

// Configuration contains the json structure for the config file
type Configuration struct {
	MongoDb MongoConfig
	App     AppConfig
	Redis   RedisConfig
}

// MongoConfig contains the fields to access and manage MongoDb connection
type MongoConfig struct {
	ConnectionString string
	Port             int
	User             string
	Password         string
	DbName           string
}

// AppConfig contains fields that have common/generic values for application startup
type AppConfig struct {
	Address string
}

// RedisConfig contains the fields to access and manage Redis connection
type RedisConfig struct {
	ConnectionString string
	User             string
	Password         string
}
