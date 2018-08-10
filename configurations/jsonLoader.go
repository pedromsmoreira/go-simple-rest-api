package configurations

import (
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

// JSONLoader - load json files with configuration values
type JSONLoader struct {
	Fs fileSystem
}

const fname string = "cfg"
const fext string = ".json"

// Load - Should Load json file given in filePath
func (cl JSONLoader) Load() (Configuration, error) {
	config := Configuration{}
	err := decodeConfigFile(cl, "dev", &config)
	if err == nil {
		return config, err
	}

	fmt.Println("Dev file not found. Will apply default file")

	err = decodeConfigFile(cl, "", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func decodeConfigFile(cl JSONLoader, env string, config *Configuration) error {
	fp := buildFilePath(cl, env)
	f, err := cl.Fs.Open(fp)
	if err == nil {
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&config)
	}

	return err
}

func buildFilePath(cl JSONLoader, env string) string {
	var filename = []string{fname, fext}

	if len(env) > 0 {
		filename = []string{fname, ".", env, fext}
	}

	dirname := cl.Fs.GetDirName()

	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}
