package server

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/m7kvqbe1/unix-sockets-play/pb"
	"google.golang.org/protobuf/proto"
)

func StartUnixSocketServer(socketPath string) {
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Error listening: %v\n", err)
	}

	defer listener.Close()

	log.Println("Server listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err.Error())
			continue
		}

		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	var data []byte

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println("Server read error:", err)
			return
		}

		data = append(data, buf[:n]...)
	}

	var msg pb.SimpleMessage

	if err := proto.Unmarshal(data, &msg); err != nil {
		log.Println("Server protobuf decode error:", err)
		return
	}

	log.Printf("Server received: %s\n", msg.Content)
}
