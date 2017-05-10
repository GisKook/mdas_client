package conn

import (
	"github.com/giskook/mdas_client/pkg"
	"github.com/giskook/mdas_client/protocol"
	"log"
)

func event_handler_server_msg_get_config(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_get_config")
	server_get_conf_pkg := p.(*protocol.ServerGetConfigPacket)
	rep_get_conf_pkg := &protocol.RepGetConfigPacket{
		Tid:        server_get_conf_pkg.Tid,
		Serial:     server_get_conf_pkg.Serial,
		SerialPort: server_get_conf_pkg.SerialPort,
		BaudRate:   115200,
	}
	c.Send(rep_get_conf_pkg.Serialize())
}
