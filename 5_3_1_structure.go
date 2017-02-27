package main

import "fmt"

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func main() {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n)
}
