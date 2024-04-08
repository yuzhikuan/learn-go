package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

/*
 * 防止竞态
 */
func readValue(key string) int {
	// 对共享资源枷锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

/*
 * 防止竞态 -- defer
 */
func readValueDefer(key string) int {
	// 对共享资源枷锁
	valueByKeyGuard.Lock()
	// defer 解锁
	defer valueByKeyGuard.Unlock()
	// 返回值
	return valueByKey[key]
}

/**
 * 使用defer关闭打开的文件
 */
func fileSize(filename string) int64 {
	// 打开文件
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 关闭文件
	defer f.Close()

	// 取文件状态信息
	info, err := f.Stat()
	if err != nil {
		// 这里defer会自动关闭文件
		return 0
	}
	size := info.Size()
	// 这里defer会自动关闭文件
	return size
}

func main() {
	valueByKey["route"] = 666
	v := readValue("route")
	fmt.Println(v)
	valueByKey["route"] = 888
	v = readValueDefer("route")
	fmt.Println(v)
	size := fileSize("main.go")
	fmt.Println(size)
}
