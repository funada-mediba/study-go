package main

import "fmt"

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func dosomething() (a int) {
	return
}

func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	q, r := div(19, 7)

	fmt.Printf("商=%d 剰余=%d\n", q, r)
	fmt.Println(dosomething())
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
	fmt.Println(plusAlias(10, 5))
	f := returnFunc()
	f()
	returnFunc()()
	callFunction(func() {
		fmt.Println("I am Super hero")
	})

	f2 := later()
	fmt.Println(f2("Golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome!"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherIns := integers()
	fmt.Println(otherIns())
}
