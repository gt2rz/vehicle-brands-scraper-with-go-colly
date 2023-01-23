package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type CarBrand struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

var Debug bool = true
var Url string = "https://www.car.info/en-se/brands"

func main() {
	start := time.Now()

	data := []CarBrand{}
	var index uint16 = 1

	c := colly.NewCollector()

	ConsoleLog("Scraping:", Url)

	c.OnHTML("a.link_grey", func(e *colly.HTMLElement) {
		data = append(data, CarBrand{
			Id:   index,
			Name: e.Text,
		})
		index++
		ConsoleLog("Added", e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		ConsoleLog("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		elapsed := time.Since(start)
		ConsoleLog("The scraper took:", elapsed)
	})

	c.Visit(Url)

	ConsoleLog("Writing to file...")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("brands.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}

	ConsoleLog("Done!")
}

func ConsoleLog(v ...interface{}) {
	if !Debug {
		return
	}
	fmt.Println(v...)
}
