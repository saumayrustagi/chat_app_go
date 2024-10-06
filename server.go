package main

import (
	"fmt"
	"syscall"
)

func main() {
	listener_sock := create_listener()
	printListenerAddress(listener_sock)
	connected_sock := accept_sock(listener_sock)
	defer closeSockets(listener_sock, connected_sock)

	textMsg := "Ok"
	if syscall.Sendto(connected_sock, []byte(textMsg), 0, nil) == nil {
	} else {
		panic("Sendto() failed")
	}
}

func closeSockets(sockets ...int) {
	for _, socket := range sockets {
		if syscall.Close(socket) != nil {
			panic(fmt.Sprintf("Close(%d) failed", socket))
		}
	}
}

func printListenerAddress(listener_sock int) {
	if temp_sock_addr, err := syscall.Getsockname(listener_sock); err == nil {
		if sa, ok := temp_sock_addr.(*syscall.SockaddrInet4); ok {
			println("Connect to Port:", sa.Port)
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
		sock_addr := syscall.SockaddrInet4{Addr: [4]byte{127, 0, 0, 1}}
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
