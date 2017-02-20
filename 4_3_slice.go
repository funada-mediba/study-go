package main

import f "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s := make([]int, 10, 20)
	f.Println(s)
	f.Println(len(s))
	f.Println(cap(s))
	s = append(s, 1)
	f.Println(s)
	copy(s, s1)
	f.Println(s)
}
