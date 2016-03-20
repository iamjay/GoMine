package packets

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Ack struct {
	PacketId     int8
	Unknown      int16
	Additional   byte
	PacketNumber []int32
}

func (p *Ack) Encode(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.BigEndian, p.PacketId); err != nil {
		return err
	}

	if err := binary.Write(buf, binary.BigEndian, p.Unknown); err != nil {
		return err
	}

	if err := binary.Write(buf, binary.BigEndian, p.Additional); err != nil {
		return err
	}

	for _, v := range p.PacketNumber {
		var temp [4]byte
		binary.BigEndian.PutUint32(temp[0:], uint32(v))
		if _, err := buf.Write(temp[1:4]); err != nil {
			return err
		}
	}

	return nil
}

func (p *Ack) Decode(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &p.PacketId); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.BigEndian, &p.Unknown); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.BigEndian, &p.Additional); err != nil {
		return err
	}

	p.PacketNumber = make([]int32, buf.Len()/3)
	for i := 0; i < 2 && buf.Len() >= 3; i++ {
		var temp [4]byte
		if _, err := io.ReadFull(buf, temp[1:4]); err != nil {
			return err
		}
		p.PacketNumber[i] = int32(binary.BigEndian.Uint32(temp[0:]))
	}

	return nil
}
