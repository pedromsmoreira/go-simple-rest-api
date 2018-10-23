package database

type Repository interface {
	//Get(key string) (string, error)
	//Set(key, value string, expiration time.Duration) error
	Ping() (string, error)
}
