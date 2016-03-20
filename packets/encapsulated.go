package packets

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	ID_ENCAPSULATED_0 = 0x80 + iota
	ID_ENCAPSULATED_1
	ID_ENCAPSULATED_2
	ID_ENCAPSULATED_3
	ID_ENCAPSULATED_4
	ID_ENCAPSULATED_5
	ID_ENCAPSULATED_6
	ID_ENCAPSULATED_7
	ID_ENCAPSULATED_8
	ID_ENCAPSULATED_9
	ID_ENCAPSULATED_A
	ID_ENCAPSULATED_B
	ID_ENCAPSULATED_C
	ID_ENCAPSULATED_D
	ID_ENCAPSULATED_E
	ID_ENCAPSULATED_F
)

type Encapsulated struct {
	PacketId int8
	Count    int32
	Payload  []byte
}

func (p *Encapsulated) Encode(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.BigEndian, p.PacketId); err != nil {
		return err
	}

	var temp [4]byte
	binary.BigEndian.PutUint32(temp[0:], uint32(p.Count))
	if _, err := buf.Write(temp[1:4]); err != nil {
		return err
	}

	_, err := buf.Write(p.Payload)
	return err
}

func (p *Encapsulated) Decode(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &p.PacketId); err != nil {
		return err
	}

	var temp [4]byte
	if _, err := io.ReadFull(buf, temp[1:4]); err != nil {
		return err
	}
	p.Count = int32(binary.BigEndian.Uint32(temp[0:]))

	p.Payload = make([]byte, buf.Len())
	_, err := io.ReadFull(buf, p.Payload)
	return err
}
