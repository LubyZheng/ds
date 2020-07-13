//二进制转十进制
package main

import (
	"ds/linearList/stack"
	"fmt"
	"math"
)

func main() {
	var num, bin int
	var stack stack.Stack
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Scan(&bin)
		stack.Push(bin)
	}
	var result int
	for i := 0; i < num; i++ {
		result += stack.Top().(int) * int(math.Pow(2, float64(i)))
		stack.Pop()
	}
	fmt.Println(result)
}
