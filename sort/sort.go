package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		flag := false
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		j := i - 1
		element := array[i]
		for ; j >= 0; j-- {
			if array[j] > element {
				array[j+1] = array[j]
			} else {
				break
			}
			array[j] = element
		}
	}
}

func ShellSort(array []int) {
	for i := len(array) / 2; i >= 1; i /= 2 { // i == gap
		for j := i; j < len(array); j++ {
			r := j - i
			element := array[j]
			for ; r >= 0; r -= i {
				if array[r] > element {
					array[r+i] = array[r]
				} else {
					break
				}
				array[r] = element
			}
		}
	}
}

func QuickSort(array []int, l, r int) {
	if l < r {
		pivot := array[l]
		head, tail := l, r
		for head < tail {
			for head < tail && pivot <= array[tail] {
				tail--
			}
			if head < tail {
				array[head] = array[tail]
				head++
			}
			for head < tail && pivot >= array[head] {
				head++
			}
			if head < tail {
				array[tail] = array[head]
				tail--
			}
		}
		array[head] = pivot
		QuickSort(array, l, head-1)
		QuickSort(array, head+1, r)
	}
}

func generateRandomArray() []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, 1000000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000000)
	}
	return array
}

func elapsedTime(f func([]int)) {
	start := time.Now()
	f(generateRandomArray())
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	elapsedTime(BubbleSort)
	elapsedTime(InsertionSort)
	elapsedTime(ShellSort)

	a := generateRandomArray()
	start := time.Now()
	QuickSort(a, 0, len(a)-1)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
