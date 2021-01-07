package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	laddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_gram_sock_cli")
	raddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_gram_sock")
	checkError(err)
	conn, err := net.DialUnix("unixgram", laddr, raddr)
	checkError(err)
	defer conn.Close()
	n, err := conn.Write([]byte("hello world!"))
	checkError(err)
	fmt.Printf("send msg n:%d\n", n)
	var msg [20]byte
	conn.Read(msg[0:])

	fmt.Println("msg is", string(msg[0:10]))
}
