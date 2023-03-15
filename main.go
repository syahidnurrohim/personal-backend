package main

import (
	"personal-backend/api/routes"
	"personal-backend/config"
	"personal-backend/tools/scrapper"
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
	go scrapper.SynchronizeJournal()
}

func forever() {
	for {
		// fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
