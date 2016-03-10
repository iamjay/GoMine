package main

import (
	"log"
	"os"
	"os/signal"

	"bitbucket.org/pathompong/gomine/server"
)

func main() {
	log.Printf("GoMine version 0.1\n")

	s := server.New()
	s.Serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <-c:
		s.Stop()
	}

	log.Printf("GoMine exited\n")
}
