package tunnel

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
)

type Packet struct {
	dataLen uint32
	data    string
}

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

	buffer := make([]byte, 4)

	_, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	data_size := binary.BigEndian.Uint32(buffer)

	buffer = make([]byte, data_size)

	_, err = conn.Read(buffer)
	if err != nil {
		return err
	}

	p := Packet{
		data_size,
		string(buffer),
	}

	for ip, conn := range servers_map {
		fmt.Println("Sending data to ip", ip)
		err = send_data(*conn, p)
		if err != nil {
			return err
		}

	}

	return nil
}

func Connect_to_clients(servers_ip []string) error {
	for _, ip := range servers_ip {
		if err := connect_to_client(ip); err != nil {
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

func send_data(conn net.Conn, p Packet) error {
	// Convert data length to bytes
	dataLenBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLenBuff, p.dataLen)
	fmt.Println(dataLenBuff)

	// Write the data length
	_, err := conn.Write(dataLenBuff)
	if err != nil {
		return err
	}

	// Write the data
	_, err = conn.Write([]byte(p.data))
	if err != nil {
		return err
	}

	fmt.Println("Sent data successfully")
	return nil
}
