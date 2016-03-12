package session

import (
	"encoding/hex"
	"log"
	"net"

	"bitbucket.org/pathompong/gomine/packets"
)

type Session struct {
	Server Server
	Conn   *net.UDPConn
	Remote *net.UDPAddr
}

func (s *Session) SendPacket(p interface{}) error {
	data, err := packets.MarshalPacket(p)
	if err != nil {
		return err
	}
	log.Printf(">>> %s\n%s\n", s.Remote.String(), hex.Dump(data))
	_, err = s.Conn.WriteToUDP(data, s.Remote)
	if err != nil {
		return err
	}

	return nil
}
