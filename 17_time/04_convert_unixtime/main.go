package main

import (
	"time"
	"fmt"
)

func main() {
	// 时间戳转换成日期字符串
	unixtime := 1793093439
	timeObj := time.Unix(int64(unixtime), 0)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str) // 2026-10-27 17:30:39
}