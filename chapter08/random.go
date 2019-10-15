package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	随机数
	*/
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num2 := rand.Intn(10)
		fmt.Println(num2)
	}

	// 每次都不同
	rand.Seed(time.Now().Unix())
	num3 := rand.Intn(10)
	fmt.Println(num3)
}
