package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"io"
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
	log.Println("User disconnected: ", c.RemoteAddr())
}

func (u *user) close() {
	u.wait <- false
}

func (u *user) read() {
	var err error
	var command string
	buf := bufio.NewReader(u.conn)
	for {
		command, err = buf.ReadString(byte(';'))
		if err != nil {
			break
		}

		log.Println("Message:", command)
	}
	if err != io.EOF {
		log.Println("Err reading websocket: ", err.Error())
	}
	u.close()
}
