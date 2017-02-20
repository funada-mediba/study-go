package main

import (
	"fmt"
	"runtime"
)

var S = ""

func init() {
	S = S + "A"
}
func init() {
	S = S + "B"
}
func init() {
	S = S + "C"
}

func runDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("DONE")
}

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println(i)
			break
		} else {
			fmt.Println("5以外だよ")
		}
	}

	if x, y := 1, 2; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		i++
	}

	fruits := [3]string{"apple", "banana", "cherry"}
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	n := 1
	switch n {
	case 0:
		fmt.Println("1")
	case 1:
		fmt.Println("2")
	default:
		fmt.Println("unknown")
	}

L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは通らない")
	}

	runDefer()

	fmt.Printf("NumCPU:%d\n", runtime.NumCPU())
	fmt.Println(S)

}
