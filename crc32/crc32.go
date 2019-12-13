package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"net"
	"os"
	"strings"
)

func main() { // 在此包中，CRC多项式以反转符号表示，
	// 或LSB优先表示。
	//// LSB优先表示是一个带有n位的十六进制数，其中
	// 最高有效位表示x⁰和最低有效系数
	// bit表示xⁿ-1的系数（xⁿ的系数是隐含的）。
	//// 例如，CRC32-Q，由以下多项式定义，
	//	x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
	// 具有反转符号0b11010101100000101000001010000001，所以该值
	// 应该传递给MakeTable的是0xD5828281。
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32q))
	addr, err := net.ResolveUDPAddr("udp", "8081")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
			continue
		}
		strData := string(data)
		log.Println("Received:", strData)
		upper := strings.ToUpper(strData)
		_, err = conn.WriteToUDP([]byte(upper), rAddr)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Send:", upper)
	}
}

//注册消息
// 0xB0 0x00 0x00 0x00 0x89 0x54 0x33 0x11 0x12
// 0xB0 0x54 0x00 0x03
