package packets

import (
	"encoding/hex"
	"testing"
)

func TestEncapsulatedEncode(t *testing.T) {
	p := Encapsulated{
		PacketId: 100,
		Count:    10000,
		Payload:  []byte{11, 22, 33, 44, 55, 66},
	}

	data, err := MarshalPacket(&p)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", hex.Dump(data))

	var q Encapsulated
	err = UnmarshalPacket(data, &q)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v\n", q)
}
