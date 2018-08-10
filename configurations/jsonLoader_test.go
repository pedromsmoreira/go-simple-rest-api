package configurations_test

import (
	"testing"

	"github.com/pedromsmoreira/go-simple-rest-api/configurations"
)

func TestJSONLoaderLoadShouldReturnNilErr(t *testing.T) {
	// mock os.Open
	loader := configurations.JSONLoader{}
	_, err := loader.Load()

	if err != nil {
		t.Fatal("Test failed because there is no cfg.json file at least.")
	}
}
