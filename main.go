package main

import (
	"fmt"
	"vehicle-brands-scrapper-with-go/scrapers"
	"vehicle-brands-scrapper-with-go/utils"
)

var Url string = "https://www.car.info/en-se/brands"

func main() {

	data := scrapers.Brands(Url)

	err := utils.SaveInJsonFile("brands", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.ConsoleLog("Done!")
}
