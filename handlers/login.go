package handlers

import (
	"bitbucket.org/pathompong/gomine/packets"
	"bitbucket.org/pathompong/gomine/session"
)

var loginHandlers = map[byte]handleFunc{
	0x01: login,
}

func init() {
	registerHandler(loginHandlers)
}

func login(sess *session.Session, buf []byte) error {
	var p packets.ConnectedPingOpenConnections
	err := packets.UnmarshalPacket(buf, &p)
	if err != nil {
		return err
	}

	return sess.SendPacket(packets.UnconnectedPingOpenConnections{
		PacketId:   packets.ID_UNCONNECTED_PING_OPEN_CONNECTIONS,
		PingId:     p.PingId,
		ServerId:   sess.Server.ServerId(),
		Identifier: "MCPE;GoMine;2 7;0.14.0;0;20",
	})
}
