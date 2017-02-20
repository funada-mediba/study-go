package main

import f "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	f.Println(m)

	m1 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}

	f.Println(m1)
	f.Println(len(m1))
	delete(m1, 2)
	f.Println(m1)

}
