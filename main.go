package main

import (
	"log"
)

func main() {
	err := loadStatic()
	if err != nil {
		log.Fatal("Failed to load static files, ", err.Error())
	}

	log.Fatal("Application not finished!")
}
