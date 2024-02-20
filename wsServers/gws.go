package wsServers

// Examples from https://github.com/lxzan/gws/blob/master/examples
// Not finished
import (
	"github.com/lxzan/gws"
)

/*
func getUpgrader()gws.Upgrader{
	return gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		CheckUtf8Enabled: true,
		Recovery:         gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})
}
*/

/*
func HandlerGWSEcho(writer http.ResponseWriter, request *http.Request){
	socket, err := socket.==
}

*/

type Handler struct {
	gws.BuiltinEventHandler
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	_ = socket.WriteMessage(message.Opcode, message.Bytes())
}
