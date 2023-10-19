package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://hypeauditor.com/top-instagram-all-russia/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("HTTP request returned status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("instagram_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	doc.Find(".row__top").Each(func(index int, rowSelection *goquery.Selection) {
		rank := rowSelection.Find(".rank span").First().Text()
		username := rowSelection.Find(".contributor__name-content").Text()
		name := rowSelection.Find(".contributor__title").Text()
		writer.Write([]string{rank, name, username})
	})

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data saved to instagram_data.csv")
}
