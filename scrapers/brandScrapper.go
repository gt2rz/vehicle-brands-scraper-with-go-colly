package scrapers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"vehicle-brands-scrapper-with-go/utils"

	"github.com/gocolly/colly"
)

type CarBrand struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

type Config struct {
	Urls ConfigBrands `json:"urls"`
}

type ConfigBrands struct {
	Brands string `json:"brands"`
}

func Brands() []CarBrand {
	start := time.Now()
	utils.ConsoleLog("Starting the brand scraper...")

	data := []CarBrand{}
	var index uint16 = 1
	configuration := Config{}

	config, err := os.ReadFile("./scrapers/config.json")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		log.Fatal(err)
	}

	json.Unmarshal(config, &configuration)
	url := configuration.Urls.Brands

	c := colly.NewCollector()

	utils.ConsoleLog("Scraping:", url)

	c.OnHTML("a.link_grey", func(e *colly.HTMLElement) {
		data = append(data, CarBrand{
			Id:   index,
			Name: e.Text,
		})
		index++
		utils.ConsoleLog("Added", e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		utils.ConsoleLog("Something went wrong: ", err)
	})

	c.OnScraped(func(r *colly.Response) {
		elapsed := time.Since(start)
		utils.ConsoleLog("The scraper took: ", elapsed)
	})

	c.Visit(url)

	return data
}
