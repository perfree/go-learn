package main

import "fmt"

func main() {
	/**
	切片 slice,同数组类似,可变长,引用类型的数据
	*/
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)

	// make创建slice,第一个参数为类型,第二个为长度,第三个为容量,它也是可变长的
	s3 := make([]int, 3, 8)
	s3[0] = 1
	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3)
}
