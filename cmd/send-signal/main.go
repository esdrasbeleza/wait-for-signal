package main

import (
	"fmt"
	"net"
	"os"
)

var signal = ""

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("Use: send-signal <signal> <host:port>")
		return
	}

	signal = arguments[1]

	serverAndPort := arguments[2]

	conn, err := net.Dial("tcp", serverAndPort)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Printf("Sending \"%s\"... ", signal)

	if _, err := fmt.Fprintf(conn, "%s\n", signal); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}
