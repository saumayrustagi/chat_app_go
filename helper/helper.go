package helper

import (
	"fmt"
	"syscall"
)

const BUFFER_SIZE = 2048

func Recv(sock int, buffer []byte) int {
	if recvInt, _, err := syscall.Recvfrom(sock, buffer, 0); err == nil {
		return recvInt
	}
	panic("RecvFrom() failed")
}

func MakeBuffer() []byte {
	return make([]byte, BUFFER_SIZE)
}

func CloseSockets(sockets ...int) {
	for _, socket := range sockets {
		if syscall.Close(socket) != nil {
			panic(fmt.Sprintf("Close(%d) failed", socket))
		}
	}
}
