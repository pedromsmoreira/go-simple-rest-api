package configurations_test

import (
	"testing"

	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
)

func TestJSONLoaderLoadShouldReturnNilErr(t *testing.T) {
	loader := configurations.JSONLoader{
		Fs: configurations.OsFS{},
	}
	_, err := loader.Load()

	if err != nil {
		t.Fatal("Test failed because there is no cfg.json file at least.")
	}
}

func TestJSONLoaderLoadShouldReturnConfigFile(t *testing.T) {
	loader := configurations.JSONLoader{
		Fs: configurations.OsFS{},
	}
	config, _ := loader.Load()

	// need to refactor this test with fakes or mocks
	if config.MongoDb.DbName != "dummystore" {
		if config.MongoDb.DbName != "userstore" {
			t.Fatalf("Test failed. Got: %s , Wanted: %s", config.MongoDb.DbName, "userstore")
		}
	}
}
