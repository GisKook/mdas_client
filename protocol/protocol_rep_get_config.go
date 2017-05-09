package protocol

import (
	"bytes"
	"github.com/giskook/mdas_client/base"
)

type RepGetConfigPacket struct {
	Tid        uint64
	Serial     uint16
	SerialPort uint8
	NodeType   uint8
	StationID  uint8
	StartBit   uint8
	EndBit     uint8
	DataBit    uint8
	CheckBit   uint8
	BaudRate   int
}

func (p *RepGetConfigPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_GET_SERIAL_STATUS, p.Tid, p.Serial)
	writer.WriteByte(p.SerialPort)
	writer.WriteByte(p.NodeType)
	writer.WriteByte(p.StationID)
	writer.WriteByte(p.StartBit)
	writer.WriteByte(p.EndBit)
	writer.WriteByte(p.DataBit)
	writer.WriteByte(p.CheckBit)
	base.WriteDWord(&writer, uint32(p.BaudRate))
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len())-1))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
