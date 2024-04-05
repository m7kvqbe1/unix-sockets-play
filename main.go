package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/m7kvqbe1/unix-sockets-play/client"
	"github.com/m7kvqbe1/unix-sockets-play/server"
)

func main() {
	socketPath := "/tmp/example.sock"

	go server.StartUnixSocketServer(socketPath)

	time.Sleep(1 * time.Second)

	client.StartUnixSocketClient(socketPath)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
