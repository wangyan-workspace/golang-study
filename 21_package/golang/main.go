package main

import (
	"fmt"
	"golang/tools"
)

func init() {
	fmt.Println("main package initialized")
}

func main() {
	tools.PrintInfo()
}
