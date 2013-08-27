package main

import (
	"code.google.com/p/go.net/websocket"
)

type user struct {
	wait chan bool
	conn *websocket.Conn
}

func handleUser(c *websocket.Conn) {
	u := new(user)
	u.wait = make(chan bool)
	<-u.wait
}
