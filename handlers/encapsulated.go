package handlers

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"bitbucket.org/pathompong/gomine/packets"
	"bitbucket.org/pathompong/gomine/session"
)

var encapsulatedHandlers = map[byte]handleFunc{
	packets.ID_ENCAPSULATED_0: encapsulated,
	packets.ID_ENCAPSULATED_1: encapsulated,
	packets.ID_ENCAPSULATED_2: encapsulated,
	packets.ID_ENCAPSULATED_3: encapsulated,
	packets.ID_ENCAPSULATED_4: encapsulated,
	packets.ID_ENCAPSULATED_5: encapsulated,
	packets.ID_ENCAPSULATED_6: encapsulated,
	packets.ID_ENCAPSULATED_7: encapsulated,
	packets.ID_ENCAPSULATED_8: encapsulated,
	packets.ID_ENCAPSULATED_9: encapsulated,
	packets.ID_ENCAPSULATED_A: encapsulated,
	packets.ID_ENCAPSULATED_B: encapsulated,
	packets.ID_ENCAPSULATED_C: encapsulated,
	packets.ID_ENCAPSULATED_D: encapsulated,
	packets.ID_ENCAPSULATED_E: encapsulated,
	packets.ID_ENCAPSULATED_F: encapsulated,
}

var dataHandlers = map[byte]handleFunc{}

func init() {
	registerHandler(encapsulatedHandlers)
}

func encapsulated(sess *session.Session, data []byte) error {
	var p packets.Encapsulated
	if err := packets.UnmarshalPacket(data, &p); err != nil {
		return err
	}

	buf := bytes.NewBuffer(p.Payload)

	var encapId int8
	if err := binary.Read(buf, binary.BigEndian, &encapId); err != nil {
		return err
	}

	var length int16
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return err
	}
	length /= 8

	if encapId == 0x40 || encapId == 0x60 {
		var dummy [3]byte
		if _, err := io.ReadFull(buf, dummy[0:]); err != nil {
			return err
		}
	}

	if encapId == 0x60 {
		var dummy [4]byte
		if _, err := io.ReadFull(buf, dummy[0:]); err != nil {
			return err
		}
	}

	encapData := buf.Bytes()
	if len(encapData) < 1 {
		return fmt.Errorf("encapsulated packet length is too short")
	}

	if handler, ok := dataHandlers[encapData[0]]; ok {
		return handler(sess, encapData)
	}

	log.Printf("Unhandled data packet from %s:\n%s\n", sess.Remote.String(),
		hex.Dump(encapData))
	return nil
}

func registerDataHandler(handlerFuncs map[byte]handleFunc) {
	for k, v := range handlerFuncs {
		dataHandlers[k] = v
	}
}
