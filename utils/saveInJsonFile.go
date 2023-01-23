package utils

import (
	"encoding/json"
	"os"
)

func SaveInJsonFile(fileName string, data any) error {
	url := "./data/" + fileName + ".json"
	ConsoleLog("Writing to file...")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(url, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
