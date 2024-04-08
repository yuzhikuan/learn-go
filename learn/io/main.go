package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// data := []byte("I am a hacker,呵哈!\r\n I like it!")
	// rd := bytes.NewReader(data) // 注意：*bytes.Reader 没有ReadBytes方法
	// r := bufio.NewReader(rd)
	// var delim byte = ','
	// var buf [20]byte

	// Read() 方法
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// ReadByte() 方法
	// c, err := r.ReadByte()
	// fmt.Println(string(c), err)

	// ReadBytes() 方法

	// line, err := r.ReadBytes(delim)
	// fmt.Println(string(line), err)

	// ReadLine() 方法
	// line, prefix, err := r.ReadLine()
	// fmt.Println(string(line), prefix, err)

	// ReadRune() 方法
	// ch, size, err := r.ReadRune()
	// fmt.Println(string(ch), size, err)

	// ReadSlice() 方法
	// line, err := r.ReadSlice(delim)
	// fmt.Println(string(line), err)
	// line, err = r.ReadSlice(delim)
	// fmt.Println(string(line), err)
	// line, err = r.ReadSlice(delim)
	// fmt.Println(string(line), err)

	// ReadString() 方法
	// line, err := r.ReadString(delim)
	// fmt.Println(line, err)

	// Buffered() 方法
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// rn := r.Buffered()
	// fmt.Println(rn)
	// n, err = r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// rn = r.Buffered()
	// fmt.Println(rn)
	// n, err = r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// Peek() 方法 读取指定字节数的数据，这些被读取的数据不会从缓冲区中清除
	// data := []byte("Go语言入门教程")
	// rd := bytes.NewReader(data)
	// r := bufio.NewReader(rd)
	// bl, err := r.Peek(8)
	// fmt.Println(string(bl), err)
	// bl, err = r.Peek(14)
	// fmt.Println(string(bl), err)
	// bl, err = r.Peek(20)
	// fmt.Println(string(bl), err)
	// bl, err = r.Peek(22) // 比缓冲区要大
	// fmt.Println(string(bl), err)

	// Available() 方法的功能是返回缓冲区中未使用的字节数
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// p := []byte("C语言中文网")
	// fmt.Println("写入前未使用的缓冲区为：", w.Available())
	// w.Write(p)
	// fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())

	// Buffered() 方法的功能是返回已写入当前缓冲区中的字节数
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// p := []byte("C语言中文网")
	// fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	// w.Write(p)
	// fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Buffered())
	// w.Flush()
	// fmt.Println("执行 Flush 方法后，写入的字节数为：", w.Buffered())

	// Flush() 方法的功能是把缓冲区中的数据写入底层的 io.Writer，并返回错误信息。如果成功写入，error 返回 nil，否则 error 返回错误原因。
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// p := []byte("C语言中文网")
	// w.Write(p)
	// fmt.Printf("未执行 Flush 缓冲区输出 %q\n", string(wr.Bytes()))
	// w.Flush()
	// fmt.Printf("执行 Flush 后缓冲区输出 %q\n", string(wr.Bytes()))

	// Write() 方法的功能是把字节切片 p 写入缓冲区，返回已写入的字节数 nn。如果 nn 小于 len(p)，则同时返回一个错误原因。
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// p := []byte("C语言中文网")
	// n, err := w.Write(p)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), n, err)

	// WriteByte() 方法的功能是写入一个字节，如果成功写入，error 返回 nil，否则 error 返回错误原因。
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// var c byte = 'G'
	// err := w.WriteByte(c)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), err)

	// WriteRune() 方法的功能是以 UTF-8 编码写入一个 Unicode 字符，返回写入的字节数和错误信息。
	// wr := bytes.NewBuffer(nil)
	// w := bufio.NewWriter(wr)
	// var r rune = 'G'
	// size, err := w.WriteRune(r)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), size, err)

	// WriteString() 方法的功能是写入一个字符串，并返回写入的字节数和错误信息。如果返回的字节数小于 len(s)，则同时返回一个错误说明原因。
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	s := "C语言中文网"
	n, err := w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)
}
