package main

import "fmt"

func main() {
	/**
	if语句
	*/
	num := 10
	if num > 10 {
		fmt.Println("大于10")
	} else if num == 10 {
		fmt.Println("等于10")
	} else {
		fmt.Println("小于10")
	}

	// 其他写法
	if num := 4; num > 0 {
		fmt.Println(num)
	}
}
