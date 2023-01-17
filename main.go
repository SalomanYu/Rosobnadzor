package main

import (
	"time"
	"Rosobnadzor/scraper"
	"Rosobnadzor/logger"
)

const PagesLimit = 26510

func main() {
	start := time.Now().Unix()
	logger.Log.Println("Program started.")
	for i:=1; i<=PagesLimit; i++ {
		logger.Log.Printf("Vuzes page: %d\n", i)
		ids := scraper.GetVuzesId(i)
		scraper.SaveVuzes(ids)
	}
	logger.Log.Printf("%d seconds passed.\n", time.Now().Unix() - start)
}
