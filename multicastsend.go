package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var service string
	var count int
	var delay int

	count = 1
	delay = 1

	if len(os.Args) > 1 {
		service = os.Args[1]
	}

	if len(os.Args) > 2 {
		count, _ = strconv.Atoi(os.Args[2])
	}

	if len(os.Args) > 3 {
		delay, _ = strconv.Atoi(os.Args[3])
	}

	if len(os.Args) > 1 {
		ping(service, count, delay)
	}

	os.Exit(0)
}

func ping(a string, count int, delay int) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		checkError(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for i := 0; i < count; i++ {
		_, err = c.Write([]byte("hello, world\n"))
		checkError(err)
		time.Sleep(time.Second * time.Duration(delay))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
