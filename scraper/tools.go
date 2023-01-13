package scraper

import (
	"time"
	"Rosobnadzor/models"
	"github.com/gocolly/colly"
)

func splitTableRow(row *colly.HTMLElement) (tableRow models.TableRow) {
	row.ForEach("td", func(i int, h *colly.HTMLElement) {
		if i == 0 {
			tableRow.Title = h.Text
		} else if i == 1{
			tableRow.Text = h.Text
		}
	})
	return
}

func getBody(url string) (body *colly.HTMLElement, err error) {
	c := colly.NewCollector()
	c.SetRequestTimeout(30 * time.Second)
	c.OnHTML("body", func(h *colly.HTMLElement) {
		body = h
	})
	err = c.Visit(url)
	return
}