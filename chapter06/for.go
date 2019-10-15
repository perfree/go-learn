package main

import "fmt"

func main() {
	/**
	for循环
	*/
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 其他写法
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// break结束循环
	for j := 1; j <= 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println(j)
	}

	// continue跳过本次循环
	for j := 1; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
}
