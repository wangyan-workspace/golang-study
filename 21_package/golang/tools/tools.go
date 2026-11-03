package tools

import (
	"fmt"
	"golang/calc"
)

func init() {
	fmt.Println("tools package initialized")
}

func PrintInfo() {
	calc.Add(10, 10)
	fmt.Println("tools package PrintInfo")
}
