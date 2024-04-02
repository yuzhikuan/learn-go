package main

import (
	"bytes"
	"fmt"
)

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer

	// 遍历可变参数切片
	for _, s := range rawList {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s.(string))
	}

	// 将连接好的字节数组转换为字符串并输出
	fmt.Println(b.String())
}

// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

func main() {
	print("Hello", " World!", " zkyucn")
}
