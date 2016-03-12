package packet

import (
	"encoding/hex"
	"testing"
)

type Packet1 struct {
	Int8Value  int8
	Int16Value int16
	Int32Value int32
	Int64Value int64
	Str        string
	Magic      Magic
}

func TestMarshalPacket(t *testing.T) {
	p := Packet1{
		Int8Value:  -10,
		Int16Value: -1000,
		Int32Value: -100000,
		Int64Value: -10000000000,
		Str:        "Hello, World",
	}

	data, err := MarshalPacket(&p)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", hex.Dump(data))

	var q Packet1
	err = UnmarshalPacket(data, &q)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v\n", q)
}
