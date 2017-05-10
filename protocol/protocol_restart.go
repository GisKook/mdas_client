package protocol

import (
	"bytes"
	"github.com/giskook/mdas_client/base"
)

type RestartPacket struct {
	Tid    uint64
	Serial uint16
}

func (p *RestartPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_RESTART, p.Tid, p.Serial)
	base.WriteDWord(&writer, 1)
	base.WriteLength(&writer)

	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len())-1))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}
