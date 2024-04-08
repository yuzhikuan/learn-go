package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("I am a hacker,呵哈!\r\n I like it!")
	rd := bytes.NewReader(data) // 注意：*bytes.Reader 没有ReadBytes方法
	r := bufio.NewReader(rd)

	// Read() 方法
	// var buf [128]byte
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// ReadByte() 方法
	// c, err := r.ReadByte()
	// fmt.Println(string(c), err)

	// ReadBytes() 方法
	// var delim byte = ','
	// line, err := r.ReadBytes(delim)
	// fmt.Println(string(line), err)

	// ReadLine() 方法
	// line, prefix, err := r.ReadLine()
	// fmt.Println(string(line), prefix, err)

	// ReadRune() 方法
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)
}
