package handlers

import (
	"encoding/hex"
	"log"

	"bitbucket.org/pathompong/gomine/session"
)

var loginHandlers = map[byte]handleFunc{
	0x01: login,
}

func init() {
	registerHandler(loginHandlers)
}

func login(sess *session.Session, buf []byte) error {
	log.Printf("login:\n%s\n", hex.Dump(buf))

	return nil
}
