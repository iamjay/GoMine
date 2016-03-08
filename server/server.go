package server

import (
    "encoding/hex"
    "log"
    "net"
)

type Server struct {
}

func (self *Server) handle(remote *net.UDPAddr, buf []byte, byteCount int) {
    log.Printf("%s\n", hex.Dump(buf))
}

func (self *Server) Serve() error {
	log.Print("GoMine server is serving")

    addr, err := net.ResolveUDPAddr("udp", ":19132")
    if err != nil {
        return err
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        return err
    }
    defer conn.Close()

    for {
        buf := make([]byte, 1024)
        byteCount, remoteAddr, err := conn.ReadFromUDP(buf)
        if err == nil {
            self.handle(remoteAddr, buf, byteCount)
        }
    }

    return nil
}
