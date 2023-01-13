package main

import (
	"fmt"
	"time"
	"Rosobnadzor/scraper"
)

const PagesLimit = 26510

func main() {
	start := time.Now().Unix()
	
	for i:=1; i<=PagesLimit; i++ {
		ids := scraper.GetVuzesId(i)
		scraper.SaveVuzes(ids)
	}
	fmt.Println(time.Now().Unix() - start, "seconds passed.")
}
