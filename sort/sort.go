package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	//fmt.Println(array)
}

func InsertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j > 0 && j < len(array); j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	//fmt.Println(array)
}

func generateRandomArray() []int {
	array := make([]int, 1000000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000000)
	}
	return array
}

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now().Second()
	BubbleSort(generateRandomArray())
	end := time.Now().Second()
	fmt.Println(end - start)
	start = time.Now().Second()
	InsertionSort(generateRandomArray())
	end = time.Now().Second()
	fmt.Println(end - start)
}
