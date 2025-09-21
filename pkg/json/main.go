package json_saver

import (
	"encoding/json"
	"os"
)

var (
	JSON_PATH = "./servers.json"
)

// JsonSaver saves the provided servers slice into a JSON file
func JsonSaver(servers any) (bool, error) {
	data, err := json.MarshalIndent(servers, "", "  ")
	if err != nil {
		return false, err
	}

	if err := os.WriteFile(JSON_PATH, data, 0644); err != nil {
		return false, err
	}

	return true, nil
}
