package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	m := make(map[string][]string)

	scrapper := colly.NewCollector()

	scrapper.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	scrapper.OnHTML("section:not(#ref1)", func(section *colly.HTMLElement) {
		stateName := section.ChildText("h2 > a")
		if len(stateName) < 1 {
			return
		}
		m[stateName] = []string{}
		section.ForEach("ul > li > div > a", func(i int, city *colly.HTMLElement) {
			m[stateName] = append(m[stateName], city.Text)
		})
	})

	scrapper.OnScraped(func(r *colly.Response) {
		for k, v := range m {
			fmt.Printf("%s: %d\n", k, len(v))
		}
	})

	scrapper.Visit("https://www.britannica.com/topic/list-of-cities-and-towns-in-the-United-States-2023068")
}
