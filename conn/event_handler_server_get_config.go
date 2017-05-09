package conn

import (
	"github.com/giskook/mdas_client/pkg"
	"github.com/giskook/mdas_client/protocol"
	"log"
)

func event_handler_server_msg_get_config(c *Conn, p pkg.Packet) {
	log.Println("event_handler_server_msg_get_config")
	rep_get_conf_pkg := p.(*protocol.RepGetConfigPacket)
	c.Send(rep_get_conf_pkg.Serialize())
}
