package tunnel

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	servers_map = make(map[string]*net.Conn)
	mapMutex    sync.Mutex
)

func Create_server(ip string) error {
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}

	fmt.Println("Started listenning on: ", ip)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		fmt.Println("Recived conn from ", conn.RemoteAddr())

		go func(conn net.Conn) {
			if err = handle_client(conn); err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}

func handle_client(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 10)

	_, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	data_size := int(buffer[0])

	buffer = make([]byte, data_size)

	_, err = conn.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Println(string(buffer))

	return nil
}

func Connect_to_clients(servers_ip []string) error {
	for _, ip := range servers_ip {
		if err := connect_to_client(ip); err != nil {
			return err
		}

		err := handle_conn(servers_map[ip])
		if err != nil {
			return err
		}
	}
	return nil
}

func connect_to_client(ip string) error {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}

	fmt.Println("Sucessfully connected to ip: ", ip)

	mapMutex.Lock()
	servers_map[ip] = &conn
	mapMutex.Unlock()

	return nil
}

func handle_conn(conn *net.Conn) error {

	byte_msg := []byte("ola")

	_, err := (*conn).Write(byte_msg)
	if err != nil {
		return err
	}
	return nil
}
