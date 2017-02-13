package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{}
	c := [5]int{1, 2, 3}
	var d [5]int
	sa := [3]string{}

	a[1] = 0

	fmt.Println(a[0])
	fmt.Printf("%d\n%d\n%d\n%d\n%s\n", a, b, c, d, sa)
}
