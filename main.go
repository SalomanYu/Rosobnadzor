package main

import (
	"Rosobnadzor/logger"
	"Rosobnadzor/scraper"
	"time"
)

const PagesLimit = 26510

func main() {
	start := time.Now().Unix()
	logger.Log.Println("Program started.")
	for i:=1; i<=PagesLimit; i++ {
		if i < 1200 {
			continue // Просто мы уже спарсили 1200 страниц и парсим дальше
		}
		logger.Log.Printf("Vuzes page: %d\n", i)
		ids := scraper.GetVuzesId(i)
		scraper.SaveVuzes(ids)
	}
		logger.Log.Printf("%d seconds passed.\n", time.Now().Unix() - start)
	}
