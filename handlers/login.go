package handlers

import (
	"bitbucket.org/pathompong/gomine/packets"
	"bitbucket.org/pathompong/gomine/session"
)

var loginHandlers = map[byte]handleFunc{
	packets.ID_CONNECTED_PING_OPEN_CONNECTIONS: login,
	packets.ID_OPEN_CONNECTION_REQUEST_1: openConnRequest1,
}

func init() {
	registerHandler(loginHandlers)
}

func login(sess *session.Session, buf []byte) error {
	var p packets.ConnectedPingOpenConnections
	if err := packets.UnmarshalPacket(buf, &p); err != nil {
		return err
	}

	return sess.SendPacket(packets.UnconnectedPingOpenConnections{
		PacketId:   packets.ID_UNCONNECTED_PING_OPEN_CONNECTIONS,
		PingId:     p.PingId,
		ServerId:   sess.Server.ServerId(),
		Identifier: "MCPE;GoMine;2 7;0.14.0;0;20",
	})
}

function openConnRequest1(sess *session.Session, buf []byte) error {
	var p packets.ID_OPEN_CONNECTION_REQUEST_1
	if err := packets.UnmarshalPacket(buf, &p); err != nil {
		return err
	}

	// TODO: Check the magic bytes.
	return sess.SendPacket(packets.OpenConnectionReply1{
		
	});
}
