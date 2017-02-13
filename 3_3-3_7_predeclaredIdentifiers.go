package main

import (
	"fmt"
	"math"
)

func main() {

	b := false
	n := 256
	by := byte(n)
	i64 := int64(n)
	u32 := uint32(n)
	b1 := byte(255)
	b2 := b1 + byte(255)

	fmt.Println(b, n, by, i64, u32, b2)

	ui1 := uint32(400000000)
	ui2 := uint32(4000000000)
	sum := ui_1 + ui_2

	fmt.Printf("%d + %d = %d\n", ui1, ui2, sum)

	fmt.Printf("unit32 max value = %d\n", math.MaxUint32)
}
