package handlers

import (
    "encoding/hex"
    "log"
    "net"
)

var loginHandlers = map[byte]handleFunc{
    0x01: login,
}

func init() {
    registerHandler(loginHandlers)
}

func login(remote *net.UDPAddr, buf []byte) error {
    log.Printf("login:\n%s\n", hex.Dump(buf))

    return nil
}
