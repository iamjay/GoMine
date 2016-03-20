package packets

import (
	"encoding/hex"
	"testing"
)

func TestAckEncode(t *testing.T) {
	p := Ack{
		PacketId:     100,
		Unknown:      2222,
		Additional:   0,
		PacketNumber: []int32{12345, 54321},
	}

	data, err := MarshalPacket(&p)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", hex.Dump(data))

	var q Ack
	err = UnmarshalPacket(data, &q)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v\n", q)
}
