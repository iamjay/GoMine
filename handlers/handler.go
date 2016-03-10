package handlers

import (
    "encoding/hex"
    "log"
    "net"
)

type handleFunc func(remote *net.UDPAddr, buf []byte) error

var handlers = map[byte]handleFunc{}

func Handle(remote *net.UDPAddr, buf []byte) error {
    if handler, ok := handlers[buf[0]]; ok {
        return handler(remote, buf)
    }

    log.Printf("%x:\n%s\n", buf[0], hex.Dump(buf))

    return nil
}

func registerHandler(handlerFuncs map[byte]handleFunc) {
    for k, v := range handlerFuncs {
        handlers[k] = v
    }
}
