package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// manages communication to url
	c := colly.NewCollector(
		colly.AllowedDomains("thehackernews.com"),
	)
	
	// 
	c.OnHTML("content section", func(e *colly.HTMLElement) {
		titles := e.Attr("h2")
		fmt.Println(titles)
	})
	c.Visit("https://thehackernews.com/")

}


