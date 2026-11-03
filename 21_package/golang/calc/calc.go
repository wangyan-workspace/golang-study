package calc

import "fmt"

func init() {
	fmt.Println("calc package initialized")
}

func Add(x, y int) int {
	return x + y
}
