package session

import (
	"encoding/hex"
	"log"
	"net"

	"bitbucket.org/pathompong/gomine/packets"
)

type ConnectionState int

const (
	Unconnected ConnectionState = iota
	OpenReply1
	OpenReply2
	Connected
)

type Session struct {
	Conn            *net.UDPConn
	Remote          *net.UDPAddr
	Server          Server
	ConnectionState ConnectionState
	ClientId        int64
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
