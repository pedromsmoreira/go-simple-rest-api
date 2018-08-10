package configurations

// Configuration contains the json structure for the config file
type Configuration struct {
	MongoDb MongoConfig // `json:"mongodb"`
	App     AppConfig   // `json:"app"`
}

// MongoConfig contains the fields to access and manage MongoDb connection
type MongoConfig struct {
	ConnectionString string // `json:"connectionString"`
	Port             int    // `json:"port"`
	User             string // `json:"user"`
	Password         string // `json:"password"`
	DbName           string // `json:"databaseName"`
}

// AppConfig contains fields that have common/generic values for application startup
type AppConfig struct {
	Port string // `json:"port"`
}
