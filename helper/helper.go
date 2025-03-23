package helper

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"unicode/utf8"
)

const BUFFER_SIZE = 2048

func Communication(connected_sock int) {
	closed := make(chan bool)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go SenderLoop(connected_sock)

	go func() {
		ReceiverLoop(connected_sock)
		closed <- true
	}()

	select {
	case <-closed:
		fmt.Print("======CONNECTION CLOSED======")
	case <-sigChan:
		fmt.Print("Closing Connection....")
	}
	fmt.Println()
}

func SenderLoop(connected_sock int) {
	send_buffer := MakeBuffer()
	for {
		n, err := syscall.Read(syscall.Stdin, send_buffer)
		if err != nil {
			panic("Read() failed")
		}
		if syscall.Sendto(connected_sock, send_buffer[:n], 0, nil) != nil {
			fmt.Println("Sendto() failed")
			break
		}
	}
}

func ReceiverLoop(connected_sock int) {
	for {
		receive_buffer := MakeBuffer()
		recvInt, _, err := syscall.Recvfrom(connected_sock, receive_buffer, 0)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		if recvInt == 0 { //connection closed
			break
		}
		receive_buffer = receive_buffer[:recvInt]
		var padding int
		switch {
		case recvInt > 99:
			padding = 3
		case recvInt > 9:
			padding = 2
		default:
			padding = 1
		}
		runeCount := utf8.RuneCount(receive_buffer)
		countDiff := (recvInt - runeCount) / 3
		for i := 0; i < max(40, 78-padding-runeCount-countDiff); i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%d: %s", recvInt, string(receive_buffer))
	}
}

func MakeBuffer() []byte {
	return make([]byte, BUFFER_SIZE)
}

func CloseSockets(sockets ...int) {
	for _, socket := range sockets {
		if err := syscall.Close(socket); err != nil {
			fmt.Println("Close() Error:", err)
		}
	}
}

func GetPortFromArgs(position int) int {
	port, err := strconv.Atoi(os.Args[position])
	if err != nil {
		panic("Invalid number")
	}
	return port
}

func GetAddrFromArgs(position int) [4]byte {
	stringIPArray := strings.Split(os.Args[position], ".")
	if len(stringIPArray) != 4 {
		fmt.Println("Not an IP Address!")
		panic("len(IP Address) != 4")
	}
	var byteIPArray [4]byte
	for ind, str := range stringIPArray {
		tmpint, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		byteIPArray[ind] = byte(tmpint)
	}
	return byteIPArray
}

func CreateSocket() int {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic("Error creating socket")
	}
	return sock
}
