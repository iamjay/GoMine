package main

import (
	"log"

    "bitbucket.org/pathompong/gomine/server"
)

func main() {
    log.Printf("GoMine version 0.1\n")

    s := server.Server{}
    s.Serve()
}
