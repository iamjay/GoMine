package server

import (
	"log"
	"math/rand"
	"net"

	"bitbucket.org/pathompong/gomine/handlers"
	"bitbucket.org/pathompong/gomine/session"
)

type Server struct {
	serverId int64
	conn     *net.UDPConn
	exited   chan bool
	sessions map[string]*session.Session
}

func New() *Server {
	return &Server{
		serverId: rand.Int63(),
		exited:   make(chan bool),
		sessions: make(map[string]*session.Session),
	}
}

func (s *Server) ServerId() int64 {
	return s.serverId
}

func (s *Server) processPacket(remote *net.UDPAddr, data []byte) {
	sess, ok := s.sessions[remote.String()]
	if !ok {
		log.Printf("Creating a new session for %s\n", remote.String())
		sess = &session.Session{
			Server: s,
			Conn:   s.conn,
			Remote: remote,
		}
		s.sessions[remote.String()] = sess
	}

	err := handlers.Handle(sess, data)
	if err != nil {
		log.Printf("Error processing packet for %s: %s\n",
			remote.String(), err.Error())
	}
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
			data := make([]byte, 1500)
			byteCount, remoteAddr, err := s.conn.ReadFromUDP(data)
			if err != nil {
				break
			}
			s.processPacket(remoteAddr, data[0:byteCount])
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
