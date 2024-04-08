package main

import "fmt"

/**
 * 斐波那契数列 递归实现
 */
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

/**
 * 产生一个指定长度的fibonacci数列
 */
func genFibonacci(n int) (res []int) {
	res = make([]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = fibonacci(i)
	}
	return res
}

/**
 * 阶乘 递归
 */
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	// 斐波那契额数列
	fmt.Println(genFibonacci(10))
	// 10的阶乘
	fmt.Println(Factorial(10))
}
