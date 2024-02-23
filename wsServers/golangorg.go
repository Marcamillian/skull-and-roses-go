// Package wsServers provides various websocket impementations
//
// Some extra information about things
package wsServers

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

func ServerTypeGoLang() string {
	return "golang server"
}

// HandlerGolangEcho is a handler that takes a websocket connection and returns the submitted message content
func HandlerGolangEcho(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// A wsGolangServerChat tracks the websocket connections to the server
//
// based on a tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM
type wsGolangServerChat struct {
	// TODO:Marc - might be wise to make this a mutex - so we can lock it if we happen to operate on it with separate operations
	conns map[*websocket.Conn]bool //
}

// NewWsGolangServerChat provides a server instance that allows for websocket communication
//
// based on a tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM
func NewWsGolangServerChat() *wsGolangServerChat {
	return &wsGolangServerChat{
		conns: make(map[*websocket.Conn]bool),
	}
}

// HandleWsGolangChat is a handler to create & keep open a websocket connection
//
// The readLoop function loops for the life of the connection and reads in messages and broadcasts the message to all connected clients
// based on a tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM
func (s *wsGolangServerChat) HandleWSGolangChat(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true // store the connection

	s.readLoop(ws)
}

// readLoop is a method that constantly loops, reading data out of incoming messages & send it to all connected clients
//
// based on a tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM
func (s *wsGolangServerChat) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	// loop for the life of the connection
	for {
		n, err := ws.Read(buf)

		// handle errors from reading the incoming message
		if err != nil {
			if err == io.EOF { // if the connection is dropped from the other end
				break
			}

			fmt.Println("read error: ", err)
			continue // if we exit the look here - we break the ws connection
		}

		// use the message contents

		msg := buf[:n]
		fmt.Println("incoming message: ", string(msg))

		// send message to all connections
		s.broadcast(msg)
	}
}

// broadcast is a method that sends a message to all connected websockets
//
// // based on a tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM
func (s *wsGolangServerChat) broadcast(b []byte) {

	// loop through the connections in server.conns
	for ws := range s.conns {
		// immediately imnvoked function
		go func(ws *websocket.Conn) {
			// try to write to the websocket connection
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}
