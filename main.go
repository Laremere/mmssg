package main

import (
	"code.google.com/p/go.net/websocket"
	"github.com/laremere/mmssg/games"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/win"
	"log"
	"net/http"
	"time"
)

func main() {
	err := loadStatic()
	if err != nil {
		log.Fatal("Failed to load static files, ", err.Error())
	}

	http.HandleFunc("/", serveStatic)
	http.Handle("/sock/", websocket.Handler(handleUser))
	log.Println("Starting Server")
	go http.ListenAndServe(":80", nil)
	go main2()
	wde.Run()
}

const refreshRate = time.Millisecond * 50

func main2() {
	w, err := wde.NewWindow(1000, 500)
	if err != nil {
		log.Fatal("Err creating window:", err.Error())
	}
	w.Show()
	log.Println("Created window")
	go eventHandler(w)

	active := game.Games["defender"]()
	for {
		goalTime := time.Now().Add(refreshRate)
		active.Update()
		log.Println("updated")
		if time.Now().Before(goalTime) {
			time.Sleep(goalTime.Sub(time.Now()))
		}
	}
}

func eventHandler(w wde.Window) {
	events := w.EventChan()
	for event := range events {
		switch event := event.(type) {
		case wde.CloseEvent:
			err := w.Close()
			if err != nil {
				log.Fatal("Error closing window: ", err.Error())
			}
		default:
			_ = event
		}
	}
}
