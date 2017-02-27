package main

import "fmt"

type Base struct {
	Id    int
	Owner string
}

type A struct {
	Base
	Name string
	Area string
}

func main() {

	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	fmt.Println(a.Id)
	fmt.Println(a.Owner)
}
