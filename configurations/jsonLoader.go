package configurations

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// JSONLoader - load json files with configuration values
type JSONLoader struct {
}

const fname string = "cfg"
const fext string = ".json"

// Load - Should Load json file given in filePath
func (cl JSONLoader) Load() (Configuration, error) {
	config := Configuration{}
	// check if dev exist, if yes return it, fallback to no env
	f, err := os.Open(buildFilePath("dev"))
	if err == nil {
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&config)
		if err == nil {
			return config, err
		}
	}

	fmt.Println("Dev file not found. Will apply default file")

	// check if normal exists
	f, err = os.Open(buildFilePath(""))
	if err != nil {
		return config, err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func buildFilePath(env string) string {
	var filename = []string{fname, fext}

	if len(env) > 0 {
		filename = []string{fname, ".", env, fext}
	}

	_, dirname, _, _ := runtime.Caller(0)

	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}
