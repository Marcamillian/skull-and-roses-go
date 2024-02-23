package wsServers

// https://gowebexamples.com/websockets/

import (
	"fmt"
	"net/http"
	"time"
)

func ServerTypeGorilla() string {
	return "gorilla server"
}

// HandlerGorillaEcho is a handler (implements the handler interface) to echo back messages on a websocket connection
//
// relies on the
func HandlerGorillaEcho(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		msgType, msg, err := conn.ReadMessage()

		if err != nil {
			return
		}

		// print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// write message back to the browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}

		time.Sleep(time.Second)

		conn.WriteMessage(msgType, msg)
	}
}
