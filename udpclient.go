package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// CheckError function
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	// flags
	if len(os.Args) == 1 {
		fmt.Println("usage: udpclient <host> <port> <packets>")
		return
	}

	if len(os.Args) == 4 {
		ServerAddr, err := net.ResolveUDPAddr("udp", os.Args[1]+":"+os.Args[2])
		fmt.Printf("Sending UDP packet to %s.\n", os.Args[1]+":"+os.Args[2])
		CheckError(err)

		LocalAddr, err := net.ResolveUDPAddr("udp", os.Args[1]+":0")
		CheckError(err)

		Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
		CheckError(err)

		defer Conn.Close()
		i := 0
		v, err := strconv.Atoi(os.Args[3])
		CheckError(err)

		for i < v {
			msg := strconv.Itoa(i)
			i++
			buf := []byte(msg)
			_, err := Conn.Write(buf)
			if err != nil {
				fmt.Println(msg, err)
			}
			time.Sleep(time.Second * 1)
		}
	}
}
