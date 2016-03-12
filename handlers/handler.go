package handlers

import (
	"encoding/hex"
	"log"

	"bitbucket.org/pathompong/gomine/session"
)

type handleFunc func(sess *session.Session, buf []byte) error

var handlers = map[byte]handleFunc{}

func Handle(sess *session.Session, buf []byte) error {
	log.Printf("<<< %s\n%s\n", sess.Remote.String(), hex.Dump(buf))

	if handler, ok := handlers[buf[0]]; ok {
		return handler(sess, buf)
	}

	log.Printf("Unhandled packet from %s:\n%s\n", sess.Remote.String(),
		hex.Dump(buf))
	return nil
}

func registerHandler(handlerFuncs map[byte]handleFunc) {
	for k, v := range handlerFuncs {
		handlers[k] = v
	}
}
