package client

import (
	"fmt"
	"net"
)

func Connect(socketPath string) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Connected to server.")

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received:", string(buffer[:n]))
}
