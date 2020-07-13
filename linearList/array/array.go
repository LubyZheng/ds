package main

import "fmt"

func main() {
	array := make([]int, 0)
	for i := 0; i < 5; i++ {
		array = append(array, i)
	}
	//删除index=2
	array_cp := array
	result := append(array[:2], array_cp[3:]...)
	fmt.Println(result)
	//往index=2添加100
	array_cp = result
	result1 := append(result[:2], append([]int{100}, array_cp[2:]...)...)
	fmt.Println(result1)
}
