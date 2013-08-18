package main

import (
	"log"
	"net/http"
)

func main() {
	err := loadStatic()
	if err != nil {
		log.Fatal("Failed to load static files, ", err.Error())
	}

	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(":80", nil)
}
