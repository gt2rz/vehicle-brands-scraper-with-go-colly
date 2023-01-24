package scrapers

import (
	"fmt"
	"time"
	"vehicle-brands-scrapper-with-go/utils"

	"github.com/gocolly/colly"
)

type CarBrand struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

func Brands() []CarBrand {
	start := time.Now()
	utils.ConsoleLog("Starting the brand scraper...")

	data := []CarBrand{}
	var index uint16 = 1

	// Get brands config to scrape from config.json
	config, err := GetConfig()
	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}

	url := config.Scraping.Brands.Url
	element := config.Scraping.Brands.Element

	// Scrape the brands
	c := colly.NewCollector()

	utils.ConsoleLog("Scraping:", url)

	c.OnHTML(element, func(e *colly.HTMLElement) {
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
