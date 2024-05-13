package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
)

// initializing a data structure to keep the scraped data
type PhoneProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs to store the data to scrape
	var phoneProducts []PhoneProduct

	// creating a new colly instance
	c := colly.NewCollector()

	// visiting the target page
	c.Visit("https://scrapeme.live/shop/")

	// scraping logic
	c.OnHTML(".product", func(e *colly.HTMLElement) {
		phoneProduct := PhoneProduct{}

		phoneProduct.url = e.ChildAttr("a", "href")
		phoneProduct.url = e.ChildAttr("img", "src")
		phoneProduct.url = e.ChildText("h2")
		phoneProduct.url = e.ChildText(".price")

		phoneProducts = append(phoneProducts, phoneProduct)
	})

	// opening the csv file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	// writing each basketball product as a CSV row
	for _, phoneProduct := range phoneProducts {
		// converting a basketballProduct to an array of strings
		record := []string{
			phoneProduct.url,
			phoneProduct.image,
			phoneProduct.name,
			phoneProduct.price,
		}

		// writing the record to the CSV file
		writer.Write(record)
	}
	defer writer.Flush()
}
