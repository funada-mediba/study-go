package main

import "fmt"

func main() {

	s := "Taro" + " " + "Yamada"
	n := 10
	n1 := 197 | 169
	n2 := 92 ^ 137
	n3 := 165 & 155
	n4 := 108 &^ 13
	n5 := 1 &^ 1
	n6 := ^uint8(13)
	n7 := ^int8(13)
	n8 := 1 << 7
	n9 := 128 >> 3
	n += 5
	fmt.Printf("%v\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n", s, n1, n2, n3, n4, n5, n6, n7)
	fmt.Printf("%d\n%d\n%d\n", n8, n9, n)
}
