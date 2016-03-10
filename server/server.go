package server

import (
	"log"
	"net"

	"bitbucket.org/pathompong/gomine/handlers"
	"bitbucket.org/pathompong/gomine/session"
)

type Server struct {
	conn     *net.UDPConn
	exited   chan bool
	sessions map[string]*session.Session
}

func New() *Server {
	return &Server{
		exited:   make(chan bool),
		sessions: make(map[string]*session.Session),
	}
}

func (s *Server) processPacket(remote *net.UDPAddr, buf []byte) {
	sess, ok := s.sessions[remote.String()]
	if !ok {
		log.Printf("Creating a new session for %s\n", remote.String())
		sess = session.New()
		s.sessions[remote.String()] = sess
	}

	handlers.Handle(sess, buf)
}

func (s *Server) Serve() error {
	log.Print("GoMine server is serving")

	addr, err := net.ResolveUDPAddr("udp", ":19132")
	if err != nil {
		return err
	}

	s.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	go func() {
		defer s.conn.Close()

		for {
			buf := make([]byte, 1500)
			byteCount, remoteAddr, err := s.conn.ReadFromUDP(buf)
			if err != nil {
				break
			}
			s.processPacket(remoteAddr, buf[0:byteCount])
		}

		s.exited <- true
	}()

	return nil
}

func (s *Server) Stop() {
	log.Printf("GoMine server is exiting")

	s.conn.Close()
	<-s.exited
}
