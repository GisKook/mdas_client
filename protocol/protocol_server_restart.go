package protocol

import ()

type ServerRestartPacket struct {
	Tid    uint64
	Serial uint16
}

func (p *ServerRestartPacket) Serialize() []byte {
	return nil
}

func ParseServerRestart(buffer []byte) *ServerRestartPacket {
	_, _, _, tid, serial := ParseHeader(buffer)

	return &ServerRestartPacket{
		Tid:    tid,
		Serial: serial,
	}
}
