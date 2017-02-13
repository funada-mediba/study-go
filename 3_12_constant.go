package main

import "fmt"

const ONE = 1

const (
	朝の挨拶 = "おはよう"
	昼の挨拶 = "こんにちは"
	夜の挨拶 = "こんばんは"
)

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func あいさつ(m string) {
	fmt.Println(m)
}

func main() {
	x, y := one()
	fmt.Printf("x=%d, y=%d\n", x, y)

	あいさつ(昼の挨拶)
}
