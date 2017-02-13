package main

import "fmt"

func main() {

	var x interface{}
	x = 1
	x = 3.14
	x = '山'
	x = "文字列"
	x = [...]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%#v", x)
}
