package main

import (
	f "fmt"
)

func appName() string {
	const AppName = "Go Application"
	var Verison = "1.0"
	return AppName + " " + Verison
}

func main() {
	f.Println(appName())
}
