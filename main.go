package main

import (
	"log"
	"sync"

	"github.com/portilho13/worker-main/api"
	"github.com/portilho13/worker-main/tunnel"
)

const SERVER_TCP string = "127.0.0.1:1234"

var SERVERS_TCP_IPS = []string{
	"127.0.0.1:2005",
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		err := tunnel.CreateServer(SERVER_TCP)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := api.Api("127.0.0.1:7773", &SERVERS_TCP_IPS)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	// go func(SERVERS_TCP_IPS []string) {
	// 	defer wg.Done()
	// 	err := tunnel.ConnectToClients(SERVERS_TCP_IPS)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// }(SERVERS_TCP_IPS)

	wg.Wait()
}
