package main

import (
	"chat_app/helper"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	connected_sock := connectToServer()
	defer helper.CloseSockets(connected_sock)

	for {
		receive_buffer := helper.MakeBuffer()
		recvInt := helper.Recv(connected_sock, receive_buffer)
		if recvInt == 0 { //connection closed
			break
		}
		print(fmt.Sprintf("%d:", recvInt), string(receive_buffer))
	}
}

func connectToServer() int {
	port := getAddrFromArgs()
	sock := createSocket()

	sockaddr := syscall.SockaddrInet4{Port: port, Addr: [4]byte{127, 0, 0, 1}}

	if syscall.Connect(sock, &sockaddr) != nil {
		panic(fmt.Sprintf("Error connecting to %v:%d", sockaddr.Addr, sockaddr.Port))
	}
	return sock
}

func getAddrFromArgs() int {
	address, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Invalid number")
	}
	return address
}

func createSocket() int {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic("Error creating socket")
	}
	return sock
}

// func accept_sock(listener_sock int) int {
// 	if sock, _, err := syscall.Accept(listener_sock); err == nil {
// 		return sock
// 	}
// 	panic("Accept() failed")
// }

// func create_listener() int {
// 	if sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0); err == nil {
// 		sock_addr := syscall.SockaddrInet4{Addr: [4]byte{127, 0, 0, 1}}
// 		if syscall.Bind(sock, &sock_addr) == nil {
// 			if syscall.Listen(sock, 1024) == nil {
// 				return sock
// 			}
// 			panic("Listen() failed")
// 		}
// 		panic("Bind() failed")
// 	}
// 	panic("Socket() failed")
// }
