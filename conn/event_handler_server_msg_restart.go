package conn

import (
	"github.com/giskook/mdas_client/pkg"
	"github.com/giskook/mdas_client/protocol"
	"log"
)

func event_handler_server_msg_restart(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_restart")
	server_restart_pkg := p.(*protocol.ServerRestartPacket)

	restart_pkg := &protocol.RestartPacket{
		Tid:    server_restart_pkg.Tid,
		Serial: server_restart_pkg.Serial,
	}

	c.Send(restart_pkg.Serialize())
}
