package main

import (
	"chat_app/helper"
	"fmt"
	"strconv"
	"syscall"
)

func main() {
	connected_sock := createAndAcceptConnection()
	defer helper.CloseSockets(connected_sock)

	go helper.SenderLoop(connected_sock)

	helper.ReceiverLoop(connected_sock)
}

func createAndAcceptConnection() int {
	listener_sock := create_listener()
	defer helper.CloseSockets(listener_sock)

	printListenerAddress(listener_sock)

	connected_sock := accept_sock(listener_sock)
	return connected_sock
}

func printListenerAddress(listener_sock int) {
	if temp_sock_addr, err := syscall.Getsockname(listener_sock); err == nil {
		if sa, ok := temp_sock_addr.(*syscall.SockaddrInet4); ok {
			var addr string
			for i := 0; i < 4; i++ {
				addr += strconv.Itoa(int(sa.Addr[i]))
				if i != 3 {
					addr += "."
				}
			}
			println(fmt.Sprintf("Listening on %s: %d", addr, sa.Port))
		} else {
			panic("Invalid port")
		}
	}
}

func accept_sock(listener_sock int) int {
	if sock, _, err := syscall.Accept(listener_sock); err == nil {
		return sock
	}
	panic("Accept() failed")
}

func create_listener() int {
	if sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0); err == nil {
		sock_addr := syscall.SockaddrInet4{Port: 8080, Addr: helper.GetAddrFromArgs(1)}
		if syscall.Bind(sock, &sock_addr) == nil {
			if syscall.Listen(sock, 1024) == nil {
				return sock
			}
			panic("Listen() failed")
		}
		panic("Bind() failed")
	}
	panic("Socket() failed")
}
