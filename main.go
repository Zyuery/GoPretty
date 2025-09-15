package main

import "fmt"

// 函数闭包，当作出参的函数，在其调用域的内部 ,维护了一个外部环境 sum
func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := add(), add()
	for i := 0; i < 10; i++ {
		fmt.Println(i, pos(i), neg(-1*i))
	}
}
