package tunnel

import (
	"fmt"
	"log"
	"net"
)

func Create_conn(ip string) error {
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
