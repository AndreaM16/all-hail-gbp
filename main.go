package main

import (
	"fmt"
	"github.com/andream16/all-hail-gbp/configuration"
	"github.com/andream16/all-hail-gbp/quandl"
	"log"
)

func main() {
	fmt.Println("Getting configuration file . . .")
	config := configuration.InitConfiguration()
	response, err := quandl.CrawlCurrencies(&config); if err != nil {
		log.Fatal(err.Error())
	}
	for i, k := range response {
		fmt.Println("set number ", i)
		for _, j := range k {
			fmt.Println(j.Date, j.Value)
		}
	}
}
