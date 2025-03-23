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

	helper.Communication(connected_sock)
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
			for i := range 4 {
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
	sock := helper.CreateSocket()
	sock_addr := syscall.SockaddrInet4{Port: 8080, Addr: helper.GetAddrFromArgs(1)}
	if err := syscall.SetsockoptInt(sock, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		fmt.Println("SO_REUSEADDR failed:", err)
	}
	if syscall.Bind(sock, &sock_addr) == nil {
		if syscall.Listen(sock, 1024) == nil {
			return sock
		}
		panic("Listen() failed")
	}
	panic("Bind() failed")
}
