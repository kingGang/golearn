package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

//准备编码的数据
type P struct {
	X, Y, Z int
	Name    string
}

//接收解码结果的结构
type Q struct {
	X, Y, G *int32
	Name    string
}

func main() {
	data := P{3, 4, 5, "CloudGeek"}
	//编码后得到buf字节切片
	buf := encode(data)
	//用于接收解码数据
	var q *Q
	//解码操作
	q = decode(buf)
	//"CloudGeek": {3,4}
	fmt.Printf("%q: {%d,%d,%v}\n", q.Name, *q.X, *q.Y, *q)
}

func encode(data interface{}) *bytes.Buffer {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(data)
	return &buf
}

func decode(data interface{}) *Q {
	d := data.(io.Reader)
	dec := gob.NewDecoder(d)
	var q Q
	dec.Decode(&q)
	return &q
}
