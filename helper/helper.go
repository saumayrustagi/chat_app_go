package helper

import (
	"fmt"
	"syscall"
)

func Recv(sock int, buffer []byte) int {
	if recvInt, _, err := syscall.Recvfrom(sock, buffer, 0); err == nil {
		return recvInt
	}
	panic("RecvFrom() failed")
}

func CloseSockets(sockets ...int) {
	for _, socket := range sockets {
		if syscall.Close(socket) != nil {
			panic(fmt.Sprintf("Close(%d) failed", socket))
		}
	}
}
