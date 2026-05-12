package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Product struct {
	Name string
	Link string
}

func main() {
	c := colly.NewCollector()

	c.OnHTML(".product-item", func(e *colly.HTMLElement) {
		product := Product{
			Name: e.ChildText(".product-title"),
			Link: e.ChildAttr("a", "href"),
		}
		fmt.Printf("Found: %s -> %s\n", product.Name, product.Link)
	})

	c.Visit("https://nodub.com/")
}
