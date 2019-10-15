package main

import "fmt"

/**
常量的使用
*/
func main() {

	// 显式定义
	const PATH string = "http://www.baidu.com"
	fmt.Println(PATH)
	//定义一组常量
	const C1, C2, C3 = 100, 3.14, "CCC"
	fmt.Println(C1, C2, C3)
	const (
		V1 = 0
		V2 = 1
		V3 = 2
	)
	fmt.Println(V1, V2, V3)

}
