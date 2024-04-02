package main

import "fmt"

// 一个初始值为value的累加器
func Accumulator(value int) func() int { // 返回一个函数
	// 返回一个闭包
	return func() int {
		value++
		return value
	}
}

// 玩家生成器
func playerGen() func(string) (string, int) {
	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func(name string) (string, int) {
		return name, hp
	}
}

func main() {

	// 创建一个累加器，初始值为1
	accumulator := Accumulator(1)

	// 累加1并打印
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 创建一个累加器，初始值为10
	accumulator2 := Accumulator(10)

	// 累加1并打印
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)

	// 创建一个玩家生成器
	generator := playerGen()

	// 创建一个玩家，并返回玩家名字和血量
	name, hp := generator("zkyucn")

	// 打印值
	fmt.Println(name, hp)
}
