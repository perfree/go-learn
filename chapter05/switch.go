package main

import "fmt"

func main() {
	/**
	switch语句
	*/
	num := 1
	switch num {
	case 1:
		fmt.Println("等于1")
		break
	case 2:
		fmt.Println("等于2")
		break
	default:
		fmt.Println("其他")
	}
}
