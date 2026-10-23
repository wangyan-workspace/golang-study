package main

import "fmt"

func main() {
	str := "big"
	byteStr := []byte(str)
	fmt.Println(byteStr)         // [98 105 103]
	fmt.Println(string(byteStr)) // big
	byteStr[0] = 'p'
	fmt.Println(string(byteStr)) // pig

	s := "你好golang"
	runeStr := []rune(s)
	fmt.Println(runeStr) // [20320 22909 103 111 108 97 110 103]
	runeStr[0] = '我'
	fmt.Println(string(runeStr)) // 我好golang
}
