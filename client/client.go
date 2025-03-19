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
	sock := helper.CreateSocket()

	sockaddr := syscall.SockaddrInet4{Addr: helper.GetAddrFromArgs(1), Port: helper.GetPortFromArgs(2)}

	if syscall.Connect(sock, &sockaddr) != nil {
		panic(fmt.Sprintf("Error connecting to %v:%d", sockaddr.Addr, sockaddr.Port))
	}
	return sock
}
