package main

import (
	"time"

	"github.com/m7kvqbe1/unix-sockets-play/client"
	"github.com/m7kvqbe1/unix-sockets-play/server"
)

func main() {
	socketPath := "/tmp/example.sock"

	go server.StartUnixSocketServer(socketPath)

	time.Sleep(1 * time.Second)

	client.StartUnixSocketClient(socketPath)
}
