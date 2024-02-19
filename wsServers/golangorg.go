package wsServers

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

func ServerTypeGoLang() string {
	return "golang server"
}

func HandlerGolangEcho(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// == tutorial from Anthony GG - https://www.youtube.com/watch?v=JuUAEYLkGbM

type Server struct {
	// TODO:Marc - might be wise to make this a mutex - so we can lock it if we happen to operate on it with separate operations
	conns map[*websocket.Conn]bool
}

func NewServerGolangChat() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

// endpoint to initiate & keep the connection open
func (s *Server) HandleWSGolangChat(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true // store the connection

	s.readLoop(ws)
}

// constantly open loop that reads data out of incoming messages
func (s *Server) readLoop(ws *websocket.Conn) {
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
		fmt.Println(string(msg))

		ws.Write([]byte("Thanks for the message"))
	}
}
