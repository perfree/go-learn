package main

import "fmt"

func main() {
	/**
	数组
	*/
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1[0])
	fmt.Println(arr1)
	// 长度
	fmt.Println(len(arr1))
	// 容量
	fmt.Println(cap(arr1))

	// 其他书写方式
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	// 数组的遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// 数组遍历range
	for i, v := range arr2 {
		fmt.Print(i)
		fmt.Print(v)
	}
}
