package main

import (
	"fmt"
	"vehicle-brands-scrapper-with-go/scrapers"
	"vehicle-brands-scrapper-with-go/utils"
)


func main() {

	data := scrapers.Brands()

	err := utils.SaveInJsonFile("brands", data)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.ConsoleLog("Done!")
}
