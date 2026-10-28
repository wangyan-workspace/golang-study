package main

import (
	"fmt"
	"time"
)

func main() {
	var str = "2026-10-27 17:30:39"
	var tmp = "2006-01-02 15:04:05"
	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
	fmt.Println(timeObj) // 2026-10-27 17:30:39 +0800 CST
	fmt.Println(timeObj.Unix()) // 1793093439
}