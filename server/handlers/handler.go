package handlers

import (
    "encoding/hex"
    "log"
    "net"
)

type handleFunc func(remote *net.UDPAddr, buf []byte, byteCount int) error

var handlers = map[byte]handleFunc{}

func Handle(remote *net.UDPAddr, buf []byte, byteCount int) error {
    if handler, ok := handlers[buf[0]]; ok {
        return handler(remote, buf, byteCount)
    }

    log.Printf("%x:\n%s\n", buf[0], hex.Dump(buf))

    return nil
}

func registerHandler(handlerFuncs map[byte]handleFunc) {
    for k, v := range handlerFuncs {
        handlers[k] = v
    }
}
