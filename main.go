package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func main() {
	err := loadStatic()
	if err != nil {
		log.Fatal("Failed to load static files, ", err.Error())
	}

	http.HandleFunc("/", serveStatic)
	http.Handle("/sock/", websocket.Handler(handleUser))
	log.Println("Starting Server")
	http.ListenAndServe(":80", nil)
}
