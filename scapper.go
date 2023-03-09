package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type scrapper struct {
	Img            string `json:"img"`
	Character      string `json:"character"`
	DevilFruit     string `json:"devilFruit"`
	DevilFruitType string `json:"devilFruitType"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("listfist.com"),
	)

	var Data []scrapper

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		item := scrapper{
			Img:            e.ChildAttr("img", "src"),
			Character:      e.ChildText("td.col-2"),
			DevilFruit:     e.ChildText("td.col-3"),
			DevilFruitType: e.ChildText("td.col-4"),
		}

		Data = append(Data, item)
	})

	if err := c.Visit("https://listfist.com/list-of-one-piece-devil-fruits"); err != nil {
		fmt.Println(err)
	}

	content, err := json.Marshal(Data)
	if err != nil {
		fmt.Println(err)
	}

	if err = os.WriteFile("data.json", content, 0644); err != nil {
		fmt.Println(err)
	}
}
