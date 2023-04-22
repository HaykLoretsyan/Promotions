package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var server *http.Server

func init() {

	err := CreateDB("PromotionsDB", os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal(err)
	}

	err = CreateRepository()
	if err != nil {
		log.Fatal(err)
	}

	server = CreateServer(":8000")
}

func main() {

	go func() {
		for {
			err := UpdateData(os.Getenv("CSV_FILE_PATH"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(30 * time.Minute)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start: %s", err.Error())
	}
}
