package main

import (
	"log"
	"sync"

	"github.com/portilho13/worker-main/tunnel"
)

const SERVER_TCP string = "127.0.0.1:1234"

var SERVERS_TCP_IPS = []string{
	"127.0.0.1:2005",
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := tunnel.Create_server(SERVER_TCP)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	go func(SERVERS_TCP_IPS []string) {
		defer wg.Done()
		err := tunnel.Connect_to_clients(SERVERS_TCP_IPS)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(SERVERS_TCP_IPS)

	wg.Wait()
}
