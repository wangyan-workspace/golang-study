package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	fmt.Println(timeObj) // 2026-10-27 17:05:33.51738 +0800 CST m=+0.000053293

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// 2026-10-27 17:5:33
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	// 2026-10-27 17:05:33
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	//注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上
}