package main

import "fmt"

// 变量的定义及使用
func main() {

	// 第一种,先定义再赋值
	var age int
	age = 22
	fmt.Printf("age的值是: %d\n", age)

	var name string
	name = "张三"
	fmt.Print("name的值是:" + name + "\n")

	// 第二种,类型判断
	var age1 = 24
	var name1 = "王二"
	fmt.Printf(name1+"的年龄是 %d\n", age1)
	fmt.Printf("name1的类型是%T,age1的类型是%T\n", name1, age1)

	// 第三种,简短定义
	sum := 100
	fmt.Println(sum)

	// 多变量声明
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n = 100, 200
	fmt.Println(m, n)

	var n1, f1, m1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, m1)

	var (
		q1 = "测试"
		q2 = 18
		q3 = 3.14
	)
	fmt.Println(q1, q2, q3)
}
