package server

import (
	"fmt"
	"net"
	"os"

	"github.com/m7kvqbe1/unix-sockets-play/pb"
	"google.golang.org/protobuf/proto"
)

func StartUnixSocketServer(socketPath string) {
	os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}

		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Server read error:", err)
		return
	}

	var msg pb.SimpleMessage

	if err := proto.Unmarshal(buf[:n], &msg); err != nil {
		fmt.Println("Server protobuf decode error:", err)
		return
	}

	fmt.Printf("Server received: %s\n", msg.Content)
}
