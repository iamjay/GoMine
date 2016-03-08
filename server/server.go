package server

import (
    "log"
    "net"

    "bitbucket.org/pathompong/gomine/server/handlers"
)

type Server struct {
}

func New() *Server {
    return &Server {
    }
}

func (s *Server) handle(remote *net.UDPAddr, buf []byte, byteCount int) {
    handlers.Handle(remote, buf, byteCount)
}

func (s *Server) Serve() error {
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
            go s.handle(remoteAddr, buf, byteCount)
        }
    }

    return nil
}
