package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"github.com/laremere/mmssg/games"
	"io"
	"log"
	"strconv"
	"strings"
)

var userIdChan chan int = make(chan int)

func genUserIds() {
	i := 0
	for {
		userIdChan <- i
		i++
	}
}

type user struct {
	wait chan bool
	id   int
	conn *websocket.Conn
}

func handleUser(c *websocket.Conn) {
	u := new(user)
	u.id = <-userIdChan
	log.Println("User ", u.id, " connected: ", c.RemoteAddr())
	u.conn = c
	u.wait = make(chan bool)

	event := new(game.UserAddEvent)
	event.Id = u.id
	userEventChan <- event

	go u.read()
	<-u.wait
	err := c.Close()
	if err != nil {
		log.Println("Err closing websocket: ", err.Error())
	}
	log.Println("User ", u.id, " disconnected: ", c.RemoteAddr())
}

func (u *user) close() {
	u.wait <- false
	event := new(game.UserDropEvent)
	event.Id = u.id
	userEventChan <- event
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
		command = command[:len(command)-1]
		commandArgs := strings.Split(command, " ")
		if commandArgs[0] == "move" {
			x, err1 := strconv.ParseFloat(commandArgs[1], 64)
			y, err2 := strconv.ParseFloat(commandArgs[2], 64)
			event := new(game.UserMoveEvent)
			event.Id = u.id
			event.X = x
			event.Y = y
			userEventChan <- event
			if err1 != nil || err2 != nil {
				log.Println("Invalid move command from client, dropping")
				break
			}
		} else {
			log.Println("Invalid command from client, dropping")
			break
		}
	}
	if err != io.EOF {
		log.Println("Err reading websocket: ", err.Error())
	}
	u.close()
}

var userEventChan chan game.UserEvent = make(chan game.UserEvent)

func userEventCollector(request chan bool, result chan []game.UserEvent) {
	out := make([]game.UserEvent, 0)
	for {
		select {
		case <-request:
			result <- out
			out = make([]game.UserEvent, 0)
		case event := <-userEventChan:
			out = append(out, event)
		}
	}
}
