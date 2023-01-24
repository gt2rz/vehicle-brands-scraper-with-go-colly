package scrapers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Scraping struct {
		Brands struct {
			Url     string `json:"url"`
			Element string `json:"element"`
		} `json:"brands"`
	} `json:"scraping"`
}

func GetConfig() (Config, error) {
	configuration := Config{}

	config, err := os.ReadFile("./scrapers/config.json")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return configuration, err
	}

	json.Unmarshal(config, &configuration)

	return configuration, nil
}
