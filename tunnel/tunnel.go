package tunnel

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type Packet struct {
	DataLen uint32
	Data    []byte
}

var (
	ServersMap = make(map[string]net.Conn)
	MapMutex   sync.Mutex
)

func CreateServer(ip string) error {
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}
	fmt.Println("Started listening on: ", ip)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		fmt.Println("Received conn from ", conn.RemoteAddr())
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	defer conn.Close()

	var dataLen uint32
	err := binary.Read(conn, binary.BigEndian, &dataLen)
	if err != nil {
		log.Println("Failed to read data length:", err)
		return
	}

	buffer := make([]byte, dataLen)
	_, err = io.ReadFull(conn, buffer)
	if err != nil {
		log.Println("Failed to read data:", err)
		return
	}

	packet := Packet{DataLen: dataLen, Data: buffer}
	MapMutex.Lock()
	defer MapMutex.Unlock()
	for ip, conn := range ServersMap {
		fmt.Println("Sending data to ip", ip)
		if err := SendData(conn, packet); err != nil {
			log.Println("Failed to send data:", err)
			return
		}
	}
}

func ConnectToClients(serversIP []string) error {
	for _, ip := range serversIP {
		if err := ConnectToClient(ip); err != nil {
			return err
		}
	}
	return nil
}

func ConnectToClient(ip string) error {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to ip: ", ip)
	MapMutex.Lock()
	ServersMap[ip] = conn
	MapMutex.Unlock()
	return nil
}

func SendData(conn net.Conn, packet Packet) error {
	if err := binary.Write(conn, binary.BigEndian, packet.DataLen); err != nil {
		return err
	}
	_, err := conn.Write(packet.Data)
	return err
}
