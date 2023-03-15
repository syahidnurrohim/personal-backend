package main

import (
	"personal-backend/api/routes"
	"personal-backend/config"
	"personal-backend/tools/scrapper"
	"personal-backend/utils"
	"time"
)

func main() {
	// Jalankan fungsi forever di akhir
	defer forever()
	// Load environment variable
	config.LoadEnv()
	config.InitPG()
	// Jalankan server api
	go routes.Routes().Run(":8080")
	// Jalankan scrapper
	utils.SetInterval(scrapper.SynchronizeJournal, time.Hour*24)
}

func forever() {
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
