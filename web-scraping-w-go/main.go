package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string
	Author string
}

func main() {
	var quotes []Quote
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote", h.Text)
	// })

	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author", h.Text)
	// })

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q := Quote{
			Quote:  quote,
			Author: author,
		}
		quotes = append(quotes, q)
		//fmt.Printf("Quote: %s\nBy %s\n\n", quote, author)
	})

	c.Visit("http://quotes.toscrape.com/random")
	fmt.Println(quotes)
}
