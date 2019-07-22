package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var signal = ""

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("Use: wait-for-signal <string> <port number>")
		return
	}

	port := ":" + arguments[2]

	l, err := net.Listen("tcp4", port)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	signal = arguments[1]

	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {

		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		received := strings.TrimSpace(string(netData))

		if received == signal {
			fmt.Printf("Signal %s received, leaving!\n", signal)
			os.Exit(0)
		}

	}

	c.Close()
}
