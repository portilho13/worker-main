package main

import (
	"log"

	"github.com/portilho13/worker-main/tunnel"
)

const SERVER_TCP string = "127.0.0.1:1234"

var SERVERS_TCP_IPS = []string{
	"127.0.0.1:2005",
	"127.0.0.1:1986",
}

func main() {
	/* 	err := tunnel.Create_server(SERVER_TCP)
	   	if err != nil {
	   		log.Fatal(err)
	   		return
	   	} */

	err := tunnel.Connect_to_clients(SERVERS_TCP_IPS)
	if err != nil {
		log.Fatal(err)
		return
	}
}
