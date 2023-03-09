package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Img            string `json:"img"`
	Character      string `json:"character"`
	DevilFruit     string `json:"devilFruit"`
	DevilFruitType string `json:"devilFruitType"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("listfist.com"),
	)

	c.OnHTML("td", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	if err := c.Visit("https://listfist.com/list-of-one-piece-devil-fruits"); err != nil {
		fmt.Println(err)
	}
}
