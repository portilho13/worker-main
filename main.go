package main

import (
	"log"

	"github.com/portilho13/worker-main/tunnel"
)

const SERVER_TCP string = "127.0.0.1:1234"

func main() {
	err := tunnel.Create_conn(SERVER_TCP)
	if err != nil {
		log.Fatal(err)
		return
	}
}
