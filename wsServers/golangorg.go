package wsServers

import (
	"io"

	"golang.org/x/net/websocket"
)

func ServerTypeGoLang() string {
	return "golang server"
}

func HandlerGolangEcho(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
