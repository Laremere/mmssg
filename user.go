package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

type user struct {
	wait chan bool
	conn *websocket.Conn
}

func handleUser(c *websocket.Conn) {
	log.Println("User connected: ", c.RemoteAddr())
	u := new(user)
	u.conn = c
	u.wait = make(chan bool)

	go u.read()
	<-u.wait
	err := c.Close()
	if err != nil {
		log.Println("Err closing websocket: ", err.Error())
	}
}

func (u *user) close() {
	u.wait <- false
}

func (u *user) read() {
	var err error
	buf := make([]byte, 50)
	for {
		var num int
		num, err = u.conn.Read(buf)
		if err != nil {
			break
		}

		log.Println("Message:", string(buf[:num]))
		_, err = u.conn.Write(buf[:num])
		if err != nil {
			break
		}
	}
	log.Println("Err reading websocket: ", err.Error())
}
