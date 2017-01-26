package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang!")
	fmt.Println("My", "name", "is", "Taro")

	fmt.Printf("数値＝%d\n", 5)

	fmt.Printf("１０進数＝%d ２進数＝%b ８進数＝%o １６進数＝%x\n", 17, 17, 17, 17)

	fmt.Printf("数値＝%v 文字列＝%v 配列＝%v\n", 5, "Golang", [...]int{1, 2, 3})

	fmt.Printf("数値＝%#v 文字列＝%#v 配列＝%#v\n", 5, "Golang", [...]int{1, 2, 3})

	fmt.Printf("数値＝%T 文字列＝%T 配列＝%T\n", 5, "Golang", [...]int{1, 2, 3})

}
