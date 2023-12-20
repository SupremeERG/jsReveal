package misc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Contains(arr []string, target string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func ReadExistingJSON(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var existingData map[string]interface{}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&existingData); err != nil {
		return nil, err
	}

	return existingData, nil
}

func WriteJSONToFile(filePath string, data map[string]interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Convert the map to JSON
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Write the JSON to the file
	fmt.Fprintln(writer, string(jsonBytes))
	writer.Flush()

	return nil
}
