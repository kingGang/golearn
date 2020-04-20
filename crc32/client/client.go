package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	conn, err := net.Dial("udp", ":8081")
	checkError(err)
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		linelen := len(line)
		n := 0
		for written := 0; written < linelen; written += n {
			var toWrite string
			if linelen-written > 1024 {
				toWrite = line[written : written+1024]
			} else {
				toWrite = line[written:]
			}
			n, err := conn.Write([]byte(toWrite))
			checkError(err)
			fmt.Println("Write:", toWrite, n)
			msg := make([]byte, 1024)
			n, err = conn.Read(msg)
			log.Println("n=", n, "linelen=", linelen, "written=", written)
			checkError(err)
			log.Println("Response:", string(msg))
		}
	}
}
