package scraper

import (
	"Rosobnadzor/models"
	"errors"
	"time"

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
	badGateway := checkBadGateway(c)
	if badGateway {
		return nil, errors.New("Server error 502")
	}
	c.SetRequestTimeout(50 * time.Second)
	c.OnHTML("body", func(h *colly.HTMLElement) {
		body = h
	})
	err = c.Visit(url)
	return
}

func checkBadGateway(c *colly.Collector) (badGateway bool) {
	c.OnError(func(r *colly.Response, err error) {
		if r.StatusCode >= 500{
			badGateway = true
			return
		}
	})
	return
}