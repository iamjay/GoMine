package handlers

import (
	"encoding/hex"
	"fmt"
	"log"

	"bitbucket.org/pathompong/gomine/session"
)

type handleFunc func(sess *session.Session, data []byte) error

var handlers = map[byte]handleFunc{}

func Handle(sess *session.Session, data []byte) error {
	log.Printf("<<< %s\n%s\n", sess.Remote.String(), hex.Dump(data))

	if len(data) < 1 {
		return fmt.Errorf("packet length is too short")
	}

	if handler, ok := handlers[data[0]]; ok {
		return handler(sess, data)
	}

	log.Printf("Unhandled packet from %s:\n%s\n", sess.Remote.String(),
		hex.Dump(data))
	return nil
}

func registerHandler(handlerFuncs map[byte]handleFunc) {
	for k, v := range handlerFuncs {
		handlers[k] = v
	}
}
