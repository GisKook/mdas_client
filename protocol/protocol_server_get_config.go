package protocol

import ()

type ServerGetConfigPacket struct {
	Tid        uint64
	Serial     uint16
	SerialPort uint8
}

func (p *ServerGetConfigPacket) Serialize() []byte {
	return nil
}

func ParseServerGetConfig(buffer []byte) *ServerGetConfigPacket {
	reader, _, _, tid, serial := ParseHeader(buffer)
	serial_port, _ := reader.ReadByte()

	return &ServerGetConfigPacket{
		Tid:        tid,
		Serial:     serial,
		SerialPort: serial_port,
	}
}
