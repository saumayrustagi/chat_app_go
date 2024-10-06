package helper

import (
	"fmt"
	"syscall"
)

func CloseSockets(sockets ...int) {
	for _, socket := range sockets {
		if syscall.Close(socket) != nil {
			panic(fmt.Sprintf("Close(%d) failed", socket))
		}
	}
}
