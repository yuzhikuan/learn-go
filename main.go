package main

import (
	"fmt"
)

func main() {
	// 测试 字符类型（byte和rune）
	var ch1 byte = 'A'
	var ch2 rune = '\u03b2'
	var ch3 int = 87
	fmt.Printf("%X - %d\n", ch1, ch1)
	fmt.Printf("%c - %d\n", ch2, ch2)
	fmt.Printf("%c - %U\n", ch3, ch3)
}
