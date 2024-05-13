package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
)

// initializing a data structure to keep the scraped data
type BasketballProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs to store the data to scrape
	var basketballProducts []BasketballProduct

	// creating a new colly instance
	c := colly.NewCollector()

	// visiting the target page
	c.Visit("https://hoopnation.co.ke/")

	// scraping logic
	c.OnHTML("div.product", func(e *colly.HTMLElement) {
		basketballProduct := BasketballProduct{}

		basketballProduct.url = e.ChildAttr("a", "href")
		basketballProduct.url = e.ChildAttr("img", "src")
		basketballProduct.url = e.ChildText("h2")
		basketballProduct.url = e.ChildText(".price")

		basketballProducts = append(basketballProducts, basketballProduct)
	})

	

}