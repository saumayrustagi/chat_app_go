package main

import (
	"chat_app/helper"
	"fmt"
	"syscall"
)

func main() {
	connected_sock := connectToServer()
	defer helper.CloseSockets(connected_sock)

	go helper.SenderLoop(connected_sock)

	helper.ReceiverLoop(connected_sock)
}

func connectToServer() int {
	sock := createSocket()

	sockaddr := syscall.SockaddrInet4{Addr: helper.GetAddrFromArgs(1), Port: helper.GetPortFromArgs(2)}

	if syscall.Connect(sock, &sockaddr) != nil {
		panic(fmt.Sprintf("Error connecting to %v:%d", sockaddr.Addr, sockaddr.Port))
	}
	return sock
}

func createSocket() int {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic("Error creating socket")
	}
	return sock
}
