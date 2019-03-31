package main

import (
	"strconv"

	"github.com/gocolly/colly"
)

type record struct {
	wins   int64
	losses int64
}

type teams map[string]record

func getTeams() teams {
	t := teams{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"),
	)
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(_ int, el *colly.HTMLElement) {
			attr := el.Attr("data-label")
			switch attr {
			case "Team":
				var wins, _ = strconv.ParseInt(e.ChildText(`[data-label="Wins"]`), 10, 32)
				var losses, _ = strconv.ParseInt(e.ChildText(`[data-label="Losses"]`), 10, 32)
				r := record{
					wins:   wins,
					losses: losses,
				}
				t[el.Text] = r
			}
		})
	})

	c.Visit("https://erikberg.com/mlb/standings")
	return t
}
