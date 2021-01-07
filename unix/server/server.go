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

func recvUnixMsg(conn *net.UnixConn) {
	for {
		var buf [20]byte
		n, caddr, err := conn.ReadFromUnix(buf[0:])
		if err != nil {
			return
		}
		fmt.Printf("From %s Meaages:%s\n", caddr.String(), string(buf[0:]))
		n, err = conn.WriteToUnix([]byte("nice to see u"), caddr)
		checkError(err)
		fmt.Println("Success send message to ", n, caddr.String(), caddr.Network())
	}
	conn.Close()
}
func main() {
	laddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_gram_sock")
	checkError(err)
	con, err := net.ListenUnixgram("unixgram", laddr)
	checkError(err)
	recvUnixMsg(con)
}
