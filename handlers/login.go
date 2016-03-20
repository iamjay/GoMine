package handlers

import (
	"fmt"

	"bitbucket.org/pathompong/gomine/packets"
	"bitbucket.org/pathompong/gomine/session"
)

var loginHandlers = map[byte]handleFunc{
	packets.ID_UNCONNECTED_PING_OPEN_CONNECTIONS: login,
	packets.ID_OPEN_CONNECTION_REQUEST_1:         openConnRequest1,
	packets.ID_OPEN_CONNECTION_REQUEST_2:         openConnRequest2,
}

var loginDataHandlers = map[byte]handleFunc{
	packets.ID_CLIENT_CONNECT: clientConnect,
}

func init() {
	registerHandler(loginHandlers)
	registerDataHandler(loginDataHandlers)
}

func login(sess *session.Session, data []byte) error {
	var p packets.ConnectedPingOpenConnections
	if err := packets.UnmarshalPacket(data, &p); err != nil {
		return err
	}

	return sess.SendPacket(packets.UnconnectedPingOpenConnections{
		PacketId:   packets.ID_UNCONNECTED_PONG_OPEN_CONNECTIONS,
		PingId:     p.PingId,
		ServerId:   sess.Server.ServerId(),
		Identifier: "MCPE;GoMine;2 7;0.14.0;0;20",
	})
}

func openConnRequest1(sess *session.Session, data []byte) error {
	var p packets.OpenConnectionRequest1
	if err := packets.UnmarshalPacket(data, &p); err != nil {
		return err
	}

	if sess.ConnectionState > session.OpenReply1 {
		// Reset this session.
	}
	sess.ConnectionState = session.OpenReply1

	// TODO: Check the magic bytes.
	return sess.SendPacket(packets.OpenConnectionReply1{
		PacketId: packets.ID_OPEN_CONNECTION_REPLY_1,
		ServerId: sess.Server.ServerId(),
		Security: 0,
		MTUSize:  int16(len(p.Payload)),
	})
}

func openConnRequest2(sess *session.Session, data []byte) error {
	var p packets.OpenConnectionRequest2
	if err := packets.UnmarshalPacket(data, &p); err != nil {
		return err
	}

	if sess.ConnectionState != session.OpenReply1 {
		return fmt.Errorf(
			"invalid state for OpenConnectionRequest2: %v",
			sess.ConnectionState)
	}
	sess.ConnectionState = session.OpenReply2

	sess.ClientId = p.ClientId

	return sess.SendPacket(packets.OpenConnectionReply2{
		PacketId:      packets.ID_OPEN_CONNECTION_REPLY_2,
		ServerId:      sess.Server.ServerId(),
		ClientUDPPort: int16(sess.Remote.Port),
		MTUSize:       p.MTUSize,
		Security:      0,
	})
}

func clientConnect(sess *session.Session, data []byte) error {
	return nil
}
