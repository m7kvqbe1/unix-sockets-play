package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/m7kvqbe1/unix-sockets-play/client"
	"github.com/m7kvqbe1/unix-sockets-play/server"
)

func main() {
	socketPath := "/tmp/example.sock"

	if _, err := os.Stat(socketPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Access to Socket %s denied\n", socketPath)
	}

	go server.StartUnixSocketServer(socketPath)

	time.Sleep(1 * time.Second)

	client.StartUnixSocketClient(socketPath)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
