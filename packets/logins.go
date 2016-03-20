package packets

const (
	ID_UNCONNECTED_PING_OPEN_CONNECTIONS = 0x01
	ID_OPEN_CONNECTION_REQUEST_1         = 0x05
	ID_OPEN_CONNECTION_REPLY_1           = 0x06
	ID_OPEN_CONNECTION_REQUEST_2         = 0x07
	ID_OPEN_CONNECTION_REPLY_2           = 0x08
	ID_INCOMPATIBLE_PROTOCOL_VERSION     = 0x1A
	ID_UNCONNECTED_PONG_OPEN_CONNECTIONS = 0x1C
)

const (
	ID_CLIENT_CONNECT = 0x09
)

type ConnectedPingOpenConnections struct {
	PacketId int8
	PingId   int64
	Magic    Magic
}

type OpenConnectionRequest1 struct {
	PacketId int8
	Magic    Magic
	Version  int8
	Payload  []byte
}

type OpenConnectionReply1 struct {
	PacketId int8
	Magic    Magic
	ServerId int64
	Security int8
	MTUSize  int16
}

type OpenConnectionRequest2 struct {
	PacketId      int8
	Magic         Magic
	Security      int8
	Cookie        int32
	ServerUDPPort int16
	MTUSize       int16
	ClientId      int64
}

type OpenConnectionReply2 struct {
	PacketId      int8
	Magic         Magic
	ServerId      int64
	ClientUDPPort int16
	MTUSize       int16
	Security      int8
}

type IncompatibleProtocolVersion struct {
	PacketId        int8
	ProtocolVersion int8
	Magic           Magic
	ServerId        int64
}

type UnconnectedPingOpenConnections struct {
	PacketId   int8
	PingId     int64
	ServerId   int64
	Magic      Magic
	Identifier string
}

type ClientConnect struct {
	PacketId  int8
	ClientId  int64
	SessionId int64
	Unknown   int8
}
