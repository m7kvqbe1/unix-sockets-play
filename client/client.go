package client

import (
	"fmt"
	"net"

	"github.com/m7kvqbe1/unix-sockets-play/pb"
	"google.golang.org/protobuf/proto"
)

func StartUnixSocketClient(socketPath string) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	msg := &pb.SimpleMessage{Content: "Hello from client"}

	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Client protobuf encode error:", err)
		return
	}

	if _, err = conn.Write(data); err != nil {
		fmt.Println("Client write error:", err)
		return
	}
}
